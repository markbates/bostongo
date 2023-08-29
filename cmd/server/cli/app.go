package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/markbates/bostongo/web"
	"github.com/markbates/iox"
)

type App struct {
	// IO to be used by the app
	iox.IO

	// Web app to be used by the app
	web.App

	// Server to be used by the app
	// If nil, a default server will be created.
	Server *http.Server

	// Port to listen on. Defaults to 3000.
	Port int

	// Env to be used by the app
	// If nil, os.Getenv will be used.
	*Env

	mu   sync.RWMutex
	quit chan struct{}
	once sync.Once
}

// snippet: main

func (a *App) Main(ctx context.Context, pwd string, args []string) error {
	if a == nil {
		return fmt.Errorf("nil app")
	}

	flags := a.flags()
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	srv, err := a.server()
	if err != nil {
		return err
	}

	sctx, cause := context.WithCancelCause(ctx)
	defer cause(nil)

	srv.BaseContext = func(_ net.Listener) context.Context {
		return sctx
	}

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(sctx, 2*time.Second)
		defer cancel()

		cause(srv.Shutdown(ctx))

		a.mu.Lock()
		defer a.mu.Unlock()

		a.once.Do(func() {
			if a.quit != nil {
				close(a.quit)
			}
			a.quit = nil
		})

	}()

	if err := srv.ListenAndServe(); err != nil {
		cause(err)
	}

	err = context.Cause(sctx)
	if err != nil && err != context.Canceled {
		return err
	}

	return nil
}

// snippet: main

func (a *App) SetIO(oi iox.IO) {
	if a == nil {
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.IO = oi
}

func (a *App) Done() <-chan struct{} {
	if a != nil && a.quit != nil {
		return a.quit
	}

	ch := make(chan struct{})

	if a == nil {
		close(ch)
		return ch
	}

	a.mu.Lock()
	a.quit = ch
	a.mu.Unlock()

	return a.quit
}

func (a *App) Print(w io.Writer) error {
	if a == nil {
		return fmt.Errorf("nil app")
	}

	if w == nil {
		w = os.Stdout
	}

	flags := a.flags()
	flags.SetOutput(w)
	flags.Usage()

	return nil
}

func (a *App) Describe() string {
	return "launches a web server"
}

func (a *App) Getenv(key string) (s string) {
	if a == nil || a.Env == nil {
		return os.Getenv(key)
	}

	return a.Env.Getenv(key)
}

func (a *App) flags() *flag.FlagSet {

	flags := flag.NewFlagSet("server", flag.ContinueOnError)

	flags.SetOutput(a.Stderr())
	flags.IntVar(&a.Port, "port", 3000, "port to listen on")

	return flags
}

func (a *App) server() (*http.Server, error) {
	if a == nil {
		return nil, fmt.Errorf("nil app")
	}

	a.mu.RLock()
	srv := a.Server
	port := a.Port
	a.mu.RUnlock()

	if srv != nil {
		if srv.Handler == nil {
			srv.Handler = a
		}
		return srv, nil
	}

	// snippet: port
	if port == 0 {
		p := a.Getenv("PORT")
		pi, _ := strconv.Atoi(p)
		if pi == 0 {
			pi = 3000
		}
		port = pi
	}
	// snippet: port

	srv = &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}
	srv.Handler = a

	a.mu.Lock()
	a.Server = srv
	a.mu.Unlock()

	return srv, nil
}

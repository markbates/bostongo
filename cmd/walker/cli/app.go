package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"sync"

	"github.com/markbates/bostongo"
	"github.com/markbates/iox"
)

// App is the main application struct for the walker command.
type App struct {
	// IO to be used by the app
	iox.IO

	// Cab is the file system to walk. Defaults to os.DirFS(pwd).
	Cab fs.FS

	mu sync.RWMutex
}

// snippet: main

// Main is the main entry point for the walker command.
func (a *App) Main(ctx context.Context, pwd string, args []string) error {
	if a == nil {
		return fmt.Errorf("nil app")
	}

	if ctx == nil {
		ctx = context.Background()
	}

	wk := bostongo.Walker{}
	flags := a.flags(&wk)

	err := flags.Parse(args)
	if err != nil {
		return err

	}

	a.mu.RLock()

	cab := a.Cab
	oi := a.IO

	a.mu.RUnlock()

	args = flags.Args()
	if len(args) > 0 {
		pwd = args[0]
	}

	if cab == nil {
		cab = os.DirFS(pwd)
	}

	sctx, cause := context.WithCancelCause(ctx)
	defer cause(nil)

	// launch as a goroutine so if it takes too
	// long we can cancel the command.
	go func() {
		err := wk.Walk(cab, oi.Stdout())
		cause(err)
	}()

	<-sctx.Done()

	err = context.Cause(sctx)
	if err != nil && err != context.Canceled {
		return err
	}

	return nil
}

// snippet: main

// SetIO sets the IO to be used by the app.
func (a *App) SetIO(oi iox.IO) {
	if a == nil {
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.IO = oi
}

// Print prints the usage for the walker command.
func (a *App) Print(w io.Writer) error {
	if a == nil {
		return fmt.Errorf("nil app")
	}

	if w == nil {
		w = a.Stdout()
	}

	flags := a.flags(&bostongo.Walker{})
	flags.SetOutput(w)
	flags.Usage()

	return nil
}

// Describe returns a description of the walker command.
func (a *App) Describe() string {
	return "walks a path and prints the files and directories"
}

func (a *App) flags(wk *bostongo.Walker) *flag.FlagSet {

	flags := flag.NewFlagSet("walker", flag.ContinueOnError)

	flags.SetOutput(a.Stderr())
	flags.BoolVar(&wk.PrintDirs, "dirs", false, "print directories")
	flags.BoolVar(&wk.SkipFiles, "skip-files", false, "skip files")

	return flags
}

package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"

	server "github.com/markbates/bostongo/cmd/server/cli"
	walker "github.com/markbates/bostongo/cmd/walker/cli"
	"github.com/markbates/iox"
)

type App struct {
	iox.IO

	Commands *Commands

	mu sync.RWMutex
}

func (a *App) Main(ctx context.Context, pwd string, args []string) error {
	if a == nil {
		return fmt.Errorf("nil app")
	}

	flags := a.flags()
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	cmds, err := a.populateCommands()
	if err != nil {
		return err
	}

	args = flags.Args()

	if len(args) == 0 {
		if e := a.Print(a.Stderr()); e != nil {
			return e
		}
		return fmt.Errorf("no command given")
	}

	// snippet: work
	cmd, ok := cmds.Find(args[0])
	if !ok {
		if e := a.Print(a.Stderr()); e != nil {
			return e
		}
		return fmt.Errorf("unknown command %q", args[0])
	}

	if sio, ok := cmd.(SettableIO); ok {
		sio.SetIO(a.IO)
	}

	return cmd.Main(ctx, pwd, args[1:])
	// snippet: work
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

	cmds, err := a.populateCommands()
	if err != nil {
		return err
	}

	cm := cmds.Map()
	if len(cm) == 0 {
		return nil
	}

	fmt.Fprintln(w, "Commands:")
	for n, cmd := range cm {
		s := "\t%s"
		if d, ok := cmd.(Describer); ok {
			s = fmt.Sprintf("%s\t%s", s, d.Describe())
		}
		s += "\n"
		fmt.Fprintf(w, s, n)
	}

	return nil
}

func (*App) PluginName() string {
	return "bostongo"
}

func (a *App) SetIO(io iox.IO) {
	if a == nil {
		return
	}

	a.IO = io
}

func (a *App) flags() *flag.FlagSet {

	flags := flag.NewFlagSet("bostongo", flag.ContinueOnError)

	flags.SetOutput(a.Stderr())

	return flags
}

func (a *App) populateCommands() (*Commands, error) {
	if a == nil {
		return nil, fmt.Errorf("nil app")
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if a.Commands == nil {
		a.Commands = &Commands{}
	}

	if _, ok := a.Commands.Find("server"); !ok {
		a.Commands.Set("server", &server.App{})
	}

	if _, ok := a.Commands.Find("walker"); !ok {
		a.Commands.Set("walker", &walker.App{})
	}

	if _, ok := a.Commands.Find("garlic"); !ok {
		a.Commands.Set("garlic", &Garlic{})
	}

	return a.Commands, nil
}

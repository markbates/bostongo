package cli

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	_ "github.com/markbates/garlic"
	"github.com/markbates/iox"
)

type Garlic struct {
	IO       iox.IO
	Force    bool
	SkipTidy bool
}

func (g *Garlic) Main(ctx context.Context, pwd string, args []string) error {
	if g == nil {
		return fmt.Errorf("nil garlic")
	}

	flags, err := g.flags()
	if err != nil {
		return err
	}

	if err = flags.Parse(args); err != nil {
		return err
	}

	fp := filepath.Join(pwd, "garlic", "bostongo")
	main := filepath.Join(fp, "main.go")

	if _, err := os.Stat(main); err == nil {
		if !g.Force {
			return fmt.Errorf("garlic already exists: %q", main)
		}
		if err := os.RemoveAll(fp); err != nil {
			return err
		}
	}

	err = os.MkdirAll(fp, 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(main)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(garlicTmpl)
	if err != nil {
		return err
	}

	fmt.Fprintln(g.IO.Stdout(), "created garlic at", main)

	if g.SkipTidy {
		return nil
	}

	cmd := exec.CommandContext(ctx, "go", "mod", "tidy", "-v")
	cmd.Dir = fp
	cmd.Stdout = g.IO.Stdout()
	cmd.Stderr = g.IO.Stderr()
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (g *Garlic) flags() (*flag.FlagSet, error) {
	if g == nil {
		return nil, fmt.Errorf("nil garlic")
	}

	flags := flag.NewFlagSet("garlic", flag.ExitOnError)
	flags.BoolVar(&g.Force, "f", false, "force the creation of the garlic")
	flags.BoolVar(&g.SkipTidy, "skip-tidy", false, "skip running `go mod tidy`")
	return flags, nil
}

func (Garlic) Describe() string {
	return "generates a garlic cmd"
}

func (g *Garlic) SetIO(io iox.IO) {
	if g == nil {
		return
	}

	g.IO = io
}

const garlicTmpl = `package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/markbates/bostongo/cmd/bostongo/cli"
)

func main() {
	fmt.Println("Hello from Garlic!")

	args := os.Args[1:]

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := &cli.App{
		Commands: &cli.Commands{},
	}

	// add your commands here
	// app.Commands.Set("<name>", <Commander>)

	err = app.Main(ctx, pwd, args)
	if err != nil {
		log.Fatal(err)
	}
}
`

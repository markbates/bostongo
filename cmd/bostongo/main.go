package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/markbates/bostongo/cmd/bostongo/cli"
	"github.com/markbates/garlic"
)

// snippet: main
func main() {
	args := os.Args[1:]

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := &cli.App{}

	clove := &garlic.Garlic{
		Name: app.PluginName(),
		Cmd:  app,
		FS:   os.DirFS(pwd),
	}

	err = clove.Main(ctx, pwd, args)
	if err != nil {
		log.Fatal(err)
	}
}

// snippet: main

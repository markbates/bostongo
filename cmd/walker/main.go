package main

// snippet: imports
import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/markbates/bostongo/cmd/walker/cli"
)

// snippet: imports

// snippet: main
func main() {
	args := os.Args[1:]

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// snippet: app
	app := cli.App{}
	err = app.Main(ctx, pwd, args)
	// snippet: app

	if err != nil {
		log.Fatal(err)
	}
}

// snippet: main

package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/markbates/bostongo/cmd/server/cli"
)

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

	app := cli.App{}
	err = app.Main(ctx, pwd, args)
	if err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}

// snippet: main

package main

import (
	"log"
	"os"
	"os/signal"
	"context"

	"github.com/thefakequake/eventide"
)

func main() {
	c := eventide.NewClient("NzQzMzc4ODI5MzA0OTIyMTEz.XzTzfA.jliJv1jrnCSUhtkq5PoDoblukyE")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()

	if err := c.Disconnect(); err != nil {
		log.Fatalf("error while disconnecting: %s", err)
	}
}

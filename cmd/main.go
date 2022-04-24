package main

import (
	"fmt"
	"log"

	"github.com/thefakequake/eventide"
)

func main() {
	c := eventide.NewClient("Bot NzQzMzc4ODI5MzA0OTIyMTEz.XzTzfA.jliJv1jrnCSUhtkq5PoDoblukyE")
	c.LogLevel = eventide.LogInfo

	c.AddHandler(func(r *eventide.Ready) {
		fmt.Printf("Logged in as %s#%s.\n", r.User.Username, r.User.Discriminator)
	})
	c.AddHandler(func(m *eventide.MessageCreate) {
		if !(m.Content == "ping") {
			return
		}
		if _, err := c.SendMessage(m.ChannelID, "pong"); err != nil {
			log.Println(err)
		}
	})

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}

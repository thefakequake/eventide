package main

import (
	"fmt"
	"log"

	"github.com/thefakequake/eventide"
	"github.com/thefakequake/eventide/discord"
)

func main() {
	c := eventide.NewClient("Bot NzQzMzc4ODI5MzA0OTIyMTEz.XzTzfA.jliJv1jrnCSUhtkq5PoDoblukyE")
	c.LogLevel = eventide.LogInfo

	c.AddHandler(func(r *discord.Ready) {
		fmt.Printf("Logged in as %s#%s.\n", r.User.Username, r.User.Discriminator)
	})

	c.AddHandler(func(g *discord.GuildCreate) {
		fmt.Println(g.Name)
	})

	c.AddHandler(func(m *eventide.MessageCreate) {
		if m.Author.ID == c.User.ID {
			return
		}
		for _, g := range c.Guilds {
			if _, err := c.SendMessage(m.ChannelID, "I am in the guild "+g.Name); err != nil {
				log.Println(err)
			}
		}

	})

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}

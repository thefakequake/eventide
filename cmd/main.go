package main

import (
	"fmt"
	"log"
	"time"

	"github.com/thefakequake/eventide"
	"github.com/thefakequake/eventide/discord"
)

func main() {
	c := eventide.NewClient(eventide.ClientConfig{
		Token:    "NzQzMzc4ODI5MzA0OTIyMTEz.XzTzfA.jliJv1jrnCSUhtkq5PoDoblukyE",
		Intents:  discord.IntentsDefault | discord.IntentGuildMembers | discord.IntentMessageContent,
		LogLevel: eventide.LogInfo,
	})

	c.AddHandler(func(r *discord.ReadyEvent) {
		fmt.Printf("Logged in as %s#%s, using API version v%d.\n", r.User.Username, r.User.Discriminator, r.Version)
	})

	c.AddHandler(func(m *discord.MessageCreateEvent) {
		if m.Content != "messagetest" {
			return
		}
		c.CreateMessage(m.ChannelID, &discord.CreateMessage{Content: m.ID})
		time.Sleep(2 * time.Second)
		c.CreateMessage(m.ChannelID, &discord.CreateMessage{Content: m.ID})
	})

	// c.AddHandler(func(m *discord.MessageCreateEvent) {
	// 	if m.Type == discord.MessageTypeThreadStarterMessage {
	// 		return
	// 	}
	// 	fmt.Println("called")
	// 	split := strings.Split(m.Content, " ")
	// 	switch len(split) {
	// 	case 1:
	// 	case 2:
	// 		if split[0] == "thread" {
	// 			fmt.Println(m.ChannelID)
	// 			_, err := c.StartThreadFromMessage(m.ChannelID, m.ID, &discord.StartThreadFromMessage{
	// 				Name: split[1],
	// 			})
	// 			if err != nil {
	// 				log.Println(err)
	// 				return
	// 			}
	// 			fmt.Println(m.ChannelID)
	// 			if _, err := c.CreateMessage(m.ChannelID, &discord.CreateMessage{
	// 				Content:          "Hi",
	// 				MessageReference: m.Reference(),
	// 			}); err != nil {
	// 				log.Println(err.(eventide.HTTPError).Request.URL)
	// 			}
	// 		}
	// 	}
	// })

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}

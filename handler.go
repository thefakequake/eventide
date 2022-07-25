package eventide

import (
	"encoding/json"
	"reflect"

	"github.com/thefakequake/eventide/discord"
)

type Handler struct {
	Callback reflect.Value
}

// Adds an event handler based on the handler's function signature
func (c *Client) AddHandler(h any) {
	v := reflect.ValueOf(h)
	t := v.Type()
	if v.Kind() != reflect.Func {
		c.log(LogError, "event handler must be a function")
		return
	} else if t.NumIn() != 1 {
		c.log(LogError, "event handler must only have one argument")
		return
	} else if t.NumOut() > 0 {
		c.log(LogError, "event handler must not return anything")
		return
	}

	eventType := t.In(0)

	c.log(LogInfo, "registered event handler for type %s", eventType.String())

	c.handlersLock.Lock()
	c.handlers[eventType] = append(c.handlers[eventType], Handler{Callback: v})
	c.handlersLock.Unlock()
}

func (c *Client) runHandlers(op discord.GatewayPayload[json.RawMessage]) {
	e, err := eventCodec.DecodeEvent(op)
	if err != nil {
		c.log(LogWarn, "failed to decode event: %s", err)
		return
	}
	v := reflect.ValueOf(e)

	c.handlersLock.RLock()
	for _, h := range c.handlers[v.Type()] {
		h.Callback.Call([]reflect.Value{v})
	}
	c.handlersLock.RUnlock()
}

// Registers built in handlers for the client's internal use
func (c *Client) registerDefaultHandlers() {
	c.AddHandler(func(r *discord.ReadyEvent) {
		c.Lock()
		c.User = r.User
		c.sessionID = r.SessionID
		c.Unlock()
	})

	c.AddHandler(func(g *discord.GuildUpdateEvent) {
		c.guildsLock.Lock()
		c.Guilds[g.ID] = g.Guild
		c.guildsLock.Unlock()
	})

	c.AddHandler(func(g *discord.GuildCreateEvent) {
		c.guildsLock.Lock()
		c.Guilds[g.ID] = g.Guild
		c.guildsLock.Unlock()
	})

	c.AddHandler(func(g *discord.GuildDeleteEvent) {
		c.guildsLock.Lock()
		delete(c.Guilds, g.ID)
		c.guildsLock.Unlock()
	})
}

package eventide

import (
	"encoding/json"
	"reflect"

	"github.com/thefakequake/eventide/discord"
)

type Handler struct {
	Callback reflect.Value
}

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

func (c *Client) runHandlers(op Op[json.RawMessage]) {
	e, err := eventCodec.DecodeEvent(op)
	if err != nil {
		c.log(LogWarn, "failed to decode event %s: %s", op.Type, err)
		return
	}
	v := reflect.ValueOf(e)

	c.handlersLock.RLock()
	for _, h := range c.handlers[v.Type()] {
		go h.Callback.Call([]reflect.Value{v})
	}
	c.handlersLock.RUnlock()
}

func (c *Client) registerHandlers() {
	c.AddHandler(func(r *discord.Ready) {
		c.Lock()
		c.User = r.User
		c.Unlock()
	})

	c.AddHandler(func(g *discord.GuildCreate) {
		c.guildsLock.Lock()
		c.Guilds[g.ID] = g.Guild
		c.guildsLock.Unlock()
	})

	c.AddHandler(func(g *discord.GuildCreate) {
		c.guildsLock.Lock()
		c.Guilds[g.ID] = g.Guild
		c.guildsLock.Unlock()
	})

	c.AddHandler(func(g *discord.GuildDelete) {
		c.guildsLock.Lock()
		delete(c.Guilds, g.ID)
		c.guildsLock.Unlock()
	})
}

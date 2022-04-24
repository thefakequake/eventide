package eventide

import (
	"encoding/json"
	"reflect"
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

	c.handlerLock.Lock()
	c.handlers[eventType] = append(c.handlers[eventType], Handler{Callback: v})
	c.handlerLock.Unlock()
}

func (c *Client) runHandlers(op Op[json.RawMessage]) {
	e, err := eventCodec.DecodeEvent(op)
	if err != nil {
		return
	}
	v := reflect.ValueOf(e)

	c.handlerLock.RLock()
	for _, h := range c.handlers[v.Type()] {
		go h.Callback.Call([]reflect.Value{v})
	}
	c.handlerLock.RUnlock()
}

func (c *Client) registerHandlers() {
	c.AddHandler(func (r *Ready) {
		c.User = r.User
	})
}

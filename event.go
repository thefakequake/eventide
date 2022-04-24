package eventide

import (
	"encoding/json"
	"fmt"
)

var eventCodec = NewEventCodec(
	new(Ready),
	new(MessageCreate),
)

type Event interface {
	Type() string
}

func NewEventCodec(events ...Event) *EventCodec {
	c := &EventCodec{events: make(map[string]Event)}
	for _, e := range events {
		c.events[e.Type()] = e
	}
	return c
}

type EventCodec struct {
	events map[string]Event
}

func (c *EventCodec) DecodeEvent(op Op[json.RawMessage]) (Event, error) {
	e, ok := c.events[op.Type]

	if !ok {
		return nil, fmt.Errorf("unknown event: %s ", e)
	}
	err := json.Unmarshal(op.Data, &e)

	return e, err
}

type Ready struct {
	Version     int    `json:"v"`
	User        *User  `json:"user"`
	Guilds      any    `json:"guilds"`
	SessionID   string `json:"session_id"`
	Shard       [2]int `json:"shard"`
	Application any    `json:"application"`
}

func (r *Ready) Type() string { return "READY" }

type MessageCreate struct {
	*Message
}

func (r *MessageCreate) Type() string { return "MESSAGE_CREATE" }

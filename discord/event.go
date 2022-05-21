package discord

import (
	"encoding/json"
	"fmt"
)

var eventCodec = NewEventCodec(
	new(Ready),
	new(MessageCreate),
	new(GuildCreate),
	new(GuildUpdate),
	new(GuildDelete),
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

func (m *MessageCreate) Type() string { return "MESSAGE_CREATE" }

type PresenceUpdate struct {
	User       *User       `json:"user"`
	GuildID    string      `json:"guild_id"`
	Status     string      `json:"status"`
	Activities []*Activity `json:"activities"`
}

func (p *PresenceUpdate) Type() string { return "PRESENCE_UPDATE" }

type GuildCreate struct {
	*Guild
}

func (g *GuildCreate) Type() string { return "GUILD_CREATE" }

type GuildUpdate struct {
	*Guild
}

func (g *GuildUpdate) Type() string { return "GUILD_UPDATE" }

type GuildDelete struct {
	*Guild
}

func (g *GuildDelete) Type() string { return "GUILD_DELETE" }

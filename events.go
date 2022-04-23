package eventide

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Event interface {
	Type() string
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

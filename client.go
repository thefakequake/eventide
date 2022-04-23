package eventide

import (
	"net/http"
	"time"
	"reflect"

	"github.com/gorilla/websocket"
)

type Client struct {
	ws           *websocket.Conn
	http         *http.Client
	lastSequence int64
	listening    chan interface{}
	handlers     map[reflect.Type][]reflect.Value
	Token        string
	User         *User
	SessionID    string
}

func NewClient(token string) *Client {
	c := &Client{
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
		handlers: make(map[reflect.Type][]reflect.Value),
		Token:        token,
		lastSequence: 0,
	}

	return c
}

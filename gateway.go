package eventide

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/gorilla/websocket"
)

type Op[T any] struct {
	Code   int    `json:"op"`
	Sequence int64  `json:"s",omitempty`
	Type     string `json:"t",omitempty`
	Data     T      `json:"d,omitempty"`
}

type OpHello struct {
	HeartbeatInterval time.Duration `json:"heartbeat_interval"`
}

type OpIdentify struct {
	Token      string             `json:"token"`
	Intents    int                `json:"intents"`
	Properties IdentifyProperties `json:"properties"`
}

type IdentifyProperties struct {
	Os      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

// Establishes a WebSocket connection with Discord
func (c *Client) Connect() error {
	var err error

	url, err := c.GetGatewayURL()
	if err != nil {
		return err
	}

	c.ws, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	var op Op[OpHello]

	if err = c.ws.ReadJSON(&op); err != nil {
		return err
	}

	if op.Code != 10 {
		return fmt.Errorf("expected opcode 10, instead received opcode %d", p.Opcode)
	}

	if err = c.identify(); err != nil {
		return err
	}
	c.listening = make(chan interface{})

	go c.heartbeat(op.Data.HeartbeatInterval)
	go c.listenEvent()

	return nil
}

func (c *Client) identify() error {
	payload := Op[OpIdentify]{
		Code: 2,
		Data: OpIdentify{
			Token:   c.Token,
			Intents: 14023,
			Properties: IdentifyProperties{
				Os:      runtime.GOOS,
				Browser: "eventide",
				Device:  "eventide",
			},
		},
	}
	err := c.ws.WriteJSON(&payload)
	return err
}

func (c *Client) heartbeat(interval time.Duration) {
	for {
		time.Sleep(interval * time.Millisecond)
		seq := &c.lastSequence

		heartbeat := Op[*int64]{
			Code: 1,
			Data:   seq,
		}
		c.ws.WriteJSON(&heartbeat)
	}
}

func (c *Client) listenEvent() {
	for {
		var p Op[json.RawMessage]
		c.ws.ReadJSON(&p)
		switch p.Code {
		case 0:
			c.lastSequence = p.Sequence
			go c.handleEvent(p.Type, p.Data)
		case 1:
			seq := &c.lastSequence
			heartbeat := Op[*int64]{
				Code: 1,
				Data:   seq,
			}
			c.ws.WriteJSON(&heartbeat)
		case 9:
			c.identify()
		}
		
		select {
		case <-c.listening:
			return
		default:
			continue
		}
	}
}

func (c *Client) Disconnect() error {
	close(c.listening)
	err := c.ws.Close()
	return err
}

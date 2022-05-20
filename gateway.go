package eventide

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gorilla/websocket"
)

type Op[T any] struct {
	Code     int    `json:"op"`
	Sequence int64  `json:"s,omitempty"`
	Type     string `json:"t,omitempty"`
	Data     T      `json:"d,omitempty"`
}

type OpHello struct {
	HeartbeatInterval time.Duration `json:"heartbeat_interval"`
}

type OpIdentify struct {
	Token      string             `json:"token"`
	Intents    Intents            `json:"intents"`
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

	if err != nil {
		return errors.New("websocket connection is already open")
	}

	c.log(LogInfo, "fetching gateway url")
	url, err := c.GetGatewayURL()
	if err != nil {
		return err
	}

	c.Lock()
	defer c.Unlock()

	c.log(LogInfo, "connecting to the gateway")
	c.ws, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			c.Disconnect()
		}
	}()

	var op Op[OpHello]

	if err = c.ws.ReadJSON(&op); err != nil {
		return err
	}

	if op.Code != 10 {
		c.log(LogWarn, "expected opcode 10 hello, instead received opcode %d", op.Code)
	} else {
		c.log(LogInfo, "received opcode 10 hello")
	}

	c.log(LogInfo, "sending identify payload")
	if err = c.identify(); err != nil {
		return err
	}
	c.listening = make(chan any)

	go c.heartbeatLoop(op.Data.HeartbeatInterval)
	go c.listenEvent()

	return nil
}

func (c *Client) identify() error {
	payload := Op[OpIdentify]{
		Code: 2,
		Data: OpIdentify{
			Token:      c.token,
			Intents:    IntentsAll,
			Properties: c.Identify,
		},
	}
	err := c.ws.WriteJSON(&payload)
	return err
}

func (c *Client) sendHeartbeat() error {
	c.RLock()
	seq := &c.lastSequence
	c.RUnlock()

	heartbeat := Op[*int64]{
		Code: 1,
		Data: seq,
	}

	c.wsLock.Lock()
	err := c.ws.WriteJSON(&heartbeat)
	c.wsLock.Unlock()

	return err
}

func (c *Client) heartbeatLoop(interval time.Duration) {
	c.log(LogInfo, "started heartbeat goroutine")
	for {
		time.Sleep(interval * time.Millisecond)
		if err := c.sendHeartbeat(); err != nil {
			c.log(LogError, "error sending heartbeat to gateway: %s", err)
		}
	}
}

func (c *Client) listenEvent() {
	c.log(LogInfo, "started event listening goroutine")
	for {
		var op Op[json.RawMessage]
		if err := c.ws.ReadJSON(&op); err != nil {
			c.log(LogWarn, "error reading websocket message: %s", err)
			c.RLock()
			if c.ws != nil {
				c.Disconnect()
			}
			c.RUnlock()
			return
		}
		c.log(LogDebug, "op: %d seq: %d t: %s d: %s", op.Code, op.Sequence, op.Type, op.Data)
		switch op.Code {
		case 0:
			c.Lock()
			c.lastSequence = op.Sequence
			c.Unlock()
			go c.runHandlers(op)
		case 1:
			c.log(LogInfo, "sending heartbeat in response to ping")
			if err := c.sendHeartbeat(); err != nil {
				c.log(LogError, "error sending heartbeat: %s", err)
			}
		case 9:
			c.log(LogInfo, "sending identify payload in response to invalid session")
			if err := c.identify(); err != nil {
				c.log(LogError, "error sending identify payload: %s", err)
			}
		case 11:
			c.log(LogDebug, "received heartbeat ack")
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
	return c.closeWebsocket(websocket.CloseNormalClosure)
}

func (c *Client) closeWebsocket(code int) error {
	var err error

	if c.listening != nil {
		c.log(LogInfo, "closing listening channel")
		close(c.listening)
		c.listening = nil
	}

	c.Lock()
	defer c.Unlock()

	if c.ws != nil {
		c.log(LogInfo, "sending closing frame")

		c.wsLock.Lock()
		err = c.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, ""))
		c.wsLock.Unlock()
		if err != nil {
			c.log(LogInfo, "error closing websocket: %s", err)
		}

		c.log(LogInfo, "closing gateway websocket")
		err = c.ws.Close()
		if err != nil {
			c.log(LogInfo, "error closing websocket: %s", err)
		}
	}

	return err
}

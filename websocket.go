package eventide

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/gorilla/websocket"
	"github.com/thefakequake/eventide/discord"
)

// Establishes a WebSocket connection with Discord
func (c *Client) Connect() error {
	var err error

	if err != nil {
		return errors.New("websocket connection is already open")
	}

	if c.gateway == "" {
		c.gateway, err = c.GetGateway()
		if err != nil {
			return fmt.Errorf("error fetching gateway url: %s", err)
		}
		c.log(LogInfo, "fetched gateway url")
	}

	defer func() {
		if err != nil {
			c.Disconnect()
		}
	}()

	c.Lock()
	defer c.Unlock()

	c.ws, _, err = websocket.DefaultDialer.Dial(c.gateway, nil)
	if err != nil {
		return err
	}

	c.log(LogInfo, "established connection with gateway")

	var payload discord.GatewayPayload[discord.Hello]

	if err = c.ws.ReadJSON(&payload); err != nil {
		return err
	}

	if payload.Op != 10 {
		c.log(LogWarn, "expected opcode 10 hello, instead received opcode %d", payload.Op)
	} else {
		c.log(LogInfo, "received opcode 10 hello")
	}

	if c.sessionID == "" {
		if err = c.identify(); err != nil {
			return fmt.Errorf("error sending identify payload: %s", err)
		}
		c.log(LogInfo, "sent identify payload")
	} else {
		resumePayload := discord.GatewayPayload[discord.Resume]{
			Op: 6,
			Data: discord.Resume{
				Token:     c.token,
				SessionID: c.sessionID,
				Seq:       c.lastSequence,
			},
		}

		if err = c.ws.WriteJSON(&resumePayload); err != nil {
			return fmt.Errorf("error sending resume payload: %s", err)
		}
		c.log(LogInfo, "sent resume payload")
	}

	t, dat, err := c.ws.ReadMessage()
	if err != nil {
		return err
	}
	firstEvent, err := c.parsePayload(t, dat)
	if err != nil {
		return err
	}

	if firstEvent.Type != "READY" && firstEvent.Type != "RESUMED" {
		c.log(LogWarn, "expected READY or RESUMED packet, instead received: %s", firstEvent.Type)
	}

	go c.runHandlers(firstEvent)
	go c.heartbeatLoop(payload.Data.HeartbeatInterval)
	go c.listenEvent()

	return nil
}

func (c *Client) identify() error {
	payload := discord.GatewayPayload[discord.Identify]{
		Op: 2,
		Data: discord.Identify{
			Token:      c.token,
			Intents:    c.intents,
			Properties: c.identifyProperties,
			Compress:   c.compress,
		},
	}
	c.wsLock.Lock()
	err := c.ws.WriteJSON(&payload)
	c.wsLock.Unlock()

	return err
}

func (c *Client) listenClose() chan int {
	closeListener := make(chan int, 1)
	c.listenerLock.Lock()
	c.closeListeners = append(c.closeListeners, closeListener)
	c.listenerLock.Unlock()

	return closeListener
}

func (c *Client) sendHeartbeat() error {
	c.RLock()
	seq := &c.lastSequence
	c.RUnlock()

	heartbeat := discord.GatewayPayload[*int64]{
		Op:   1,
		Data: seq,
	}

	c.wsLock.Lock()
	err := c.ws.WriteJSON(&heartbeat)
	c.wsLock.Unlock()

	c.log(LogDebug, "sent heartbeat")

	return err
}

func (c *Client) heartbeatLoop(interval time.Duration) {
	c.log(LogInfo, "started heartbeat goroutine")
	listening := c.listenClose()

	ticker := time.NewTicker(interval * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		case <-listening:
			return
		}
		if err := c.sendHeartbeat(); err != nil {
			c.log(LogError, "error sending heartbeat to gateway: %s", err)
		}
	}

}

func (c *Client) listenEvent() {
	ws := c.ws
	listening := c.listenClose()

	c.log(LogInfo, "started event listening goroutine")
	for {
		t, dat, err := ws.ReadMessage()
		if err != nil {
			// check if connection wasn't closed manually by checking if connection is the same
			c.RLock()
			sameConn := ws == c.ws
			c.RUnlock()

			if sameConn {
				c.log(LogWarn, "error reading websocket message: %s", err)
				c.Reconnect()
			}

			return
		}

		payload, err := c.parsePayload(t, dat)
		if err != nil {
			c.log(LogError, "%s", err)
			continue
		}

		c.log(LogDebug, "op: %d seq: %d t: %s d: %s\n", payload.Op, payload.Sequence, payload.Type, payload.Data)
		switch payload.Op {
		case 0:
			c.Lock()
			c.lastSequence = payload.Sequence
			c.Unlock()
			go c.runHandlers(payload)
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
		case <-listening:
			return
		default:
			continue
		}
	}
}

func (c *Client) parsePayload(messageType int, dat []byte) (discord.GatewayPayload[json.RawMessage], error) {
	var reader io.Reader
	var payload discord.GatewayPayload[json.RawMessage]

	reader = bytes.NewBuffer(dat)

	if messageType == websocket.BinaryMessage {
		res, err := zlib.NewReader(reader)
		if err != nil {
			return payload, fmt.Errorf("error decompressing gateway event: %s", err)
		}

		defer func() {
			if err := res.Close(); err != nil {
				c.log(LogError, "error closing zlib: %s", err)
			}
		}()

		reader = res
	}

	if err := json.NewDecoder(reader).Decode(&payload); err != nil {
		return payload, fmt.Errorf("error decoding event JSON: %s", err)
	}

	return payload, nil
}

func (c *Client) Disconnect() error {
	return c.closeWebsocket(websocket.CloseNormalClosure)
}

func (c *Client) Reconnect() error {
	if err := c.closeWebsocket(websocket.CloseServiceRestart); err != nil {
		return err
	}
	return c.Connect()
}

func (c *Client) closeWebsocket(code int) error {
	var err error

	c.Lock()
	defer c.Unlock()

	c.listenerLock.Lock()
	closeListeners := c.closeListeners
	c.closeListeners = []chan int{}
	c.listenerLock.Unlock()

	for _, l := range closeListeners {
		l <- code
	}

	if c.ws != nil {
		c.log(LogInfo, "sending closing frame")

		c.wsLock.Lock()
		c.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, ""))
		c.wsLock.Unlock()

		c.log(LogInfo, "closing gateway websocket")
		err = c.ws.Close()
		if err != nil {
			c.log(LogInfo, "error closing websocket: %s", err)
		}
		c.ws = nil
	}

	return err
}

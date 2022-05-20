package eventide

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	sync.RWMutex

	ws           *websocket.Conn
	wsLock       sync.RWMutex
	http         *http.Client
	lastSequence int64
	listening    chan any
	handlers     map[reflect.Type][]Handler
	handlersLock sync.RWMutex
	token        string
	Identify     IdentifyProperties
	LogLevel     int
	User         *User
	SessionID    string
	Guilds       map[string]*Guild
	guildsLock   sync.RWMutex
}

func NewClient(token string) *Client {
	if !strings.HasPrefix(token, "Bot ") {
		token = "Bot " + token
	}

	c := &Client{
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
		handlers: make(map[reflect.Type][]Handler),
		token:    token,
		Identify: IdentifyProperties{
			Os:      runtime.GOOS,
			Browser: "go-eventide",
			Device:  "go-eventide",
		},
		lastSequence: 0,
	}

	c.registerHandlers()

	return c
}

func (c *Client) Run() error {
	if err := c.Connect(); err != nil {
		return fmt.Errorf("error conneting to gateway: %s", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	<-sc

	if err := c.Disconnect(); err != nil {
		return fmt.Errorf("error disconnecting from gateway: %s", err)
	}

	return nil
}

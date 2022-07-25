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
	"github.com/thefakequake/eventide/discord"
)

type Client struct {
	sync.RWMutex

	ws             *websocket.Conn
	wsLock         sync.RWMutex
	http           *http.Client
	lastSequence   int64
	closeListeners []chan int
	listenerLock   sync.RWMutex
	gateway        string
	sessionID      string

	handlers     map[reflect.Type][]Handler
	handlersLock sync.RWMutex

	token              string
	logLevel           LogLevel
	identifyProperties *discord.IdentifyConnectionProperties
	intents            discord.Intents
	compress           bool

	User       *discord.User
	Guilds     map[string]*discord.Guild
	guildsLock sync.RWMutex
}

// Client configuration
type ClientConfig struct {
	// Discord bot token
	Token string

	// Gateway intents that dictate what events the client will receive, defaults to discord.IntentsDefault
	Intents discord.Intents

	// If enabled, disables zlib data compression over the Discord Gateway
	DisableCompression bool

	// Logging level of the client
	LogLevel LogLevel

	// Gateway identify properties
	IdentifyProperties *discord.IdentifyConnectionProperties
}

func NewClient(cfg ClientConfig) *Client {
	if !strings.HasPrefix(cfg.Token, "Bot ") {
		cfg.Token = "Bot " + cfg.Token
	}
	if cfg.IdentifyProperties == nil {
		cfg.IdentifyProperties = &discord.IdentifyConnectionProperties{
			OS:      runtime.GOOS,
			Browser: "go-eventide",
			Device:  "go-eventide",
		}
	}

	c := &Client{
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
		handlers:     make(map[reflect.Type][]Handler),
		lastSequence: 0,

		token:              cfg.Token,
		logLevel:           cfg.LogLevel,
		identifyProperties: cfg.IdentifyProperties,
		intents:            cfg.Intents,
		compress:           !cfg.DisableCompression,

		Guilds: map[string]*discord.Guild{},
	}

	c.registerDefaultHandlers()

	return c
}

func (c *Client) Run() error {
	if err := c.Connect(); err != nil {
		return fmt.Errorf("error conneting to gateway: %s", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)

	listening := c.listenClose()
Loop:
	for {
		select {
		case <-sc:
			if err := c.Disconnect(); err != nil {
				return fmt.Errorf("error disconnecting from gateway: %s", err)
			}
		case code := <-listening:
			if code != websocket.CloseServiceRestart {
				break Loop
			}
			listening = c.listenClose()
		}
	}

	return nil
}

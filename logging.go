package eventide

import (
	"fmt"
	"runtime"
	"strings"
)

type LogLevel int

const (
	LogError LogLevel = iota
	LogWarn
	LogInfo
	LogDebug
)

var logTitles = map[LogLevel]string{
	LogError: "ERROR",
	LogWarn:  "WARN",
	LogInfo:  "INFO",
	LogDebug: "DEBUG",
}

var Logger = func(level LogLevel, message string, a ...any) {
	title, ok := logTitles[level]
	if ok {
		title = " " + title
	}

	_, file, line, _ := runtime.Caller(2)
	fileSplit := strings.Split(file, "/")

	fmt.Printf("[eventide%s] %s:%d: %s\n", title, fileSplit[len(fileSplit)-1], line, fmt.Sprintf(message, a...))
}

func (c *Client) log(level LogLevel, message string, a ...any) {
	if level > c.logLevel {
		return
	}
	Logger(level, message, a...)
}

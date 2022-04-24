package eventide

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	LogError int = iota
	LogWarn
	LogInfo
	LogDebug
)

var logTitles = map[int]string{
	LogError: "ERROR",
	LogWarn:  "WARN",
	LogInfo:  "INFO",
	LogDebug: "DEBUG",
}

var Logger = func(level int, message string, a ...any) {
	title, ok := logTitles[level]
	if ok {
		title = " " + title
	}

	_, file, line, _ := runtime.Caller(2)
	fileSplit := strings.Split(file, "/")

	fmt.Printf("[eventide%s] %s:%d: %s\n", title, fileSplit[len(fileSplit)-1], line, fmt.Sprintf(message, a...))
}

func (c *Client) log(level int, message string, a ...any) {
	if level > c.LogLevel {
		return
	}
	Logger(level, message, a...)
}

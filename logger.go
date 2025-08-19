package logger

import (
	"encoding/json"
	"fmt"
)

type Level int

const (
	Info Level = iota
	Warn
	Error
)

func (l Level) String() string {
	switch l {
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	default:
		return "info"
	}
}

func parseLevel(s string) Level {
	switch s {
	case "warn":
		return Warn
	case "error":
		return Error
	default:
		return Info
	}
}

type Logger struct {
	level Level
	json  bool
}

func New(level, fmtType string) *Logger {
	return &Logger{level: parseLevel(level), json: fmtType == "json"}
}

func (l *Logger) log(level Level, msg string) {
	if level < l.level {
		return
	}
	if l.json {
		entry := map[string]string{"level": level.String(), "msg": msg}
		b, _ := json.Marshal(entry)
		fmt.Println(string(b))
	} else {
		fmt.Printf("%s: %s\n", level.String(), msg)
	}
}

func (l *Logger) Info(msg string)  { l.log(Info, msg) }
func (l *Logger) Warn(msg string)  { l.log(Warn, msg) }
func (l *Logger) Error(msg string) { l.log(Error, msg) }

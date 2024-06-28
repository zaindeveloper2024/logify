package logify

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Level uint

const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
)

type Logify struct {
	out   io.Writer
	level Level
}

func New() *Logify {
	return &Logify{
		out:   os.Stdout,
		level: Info,
	}
}

func levelToString(level Level) string {
	switch level {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}
	return ""
}

func outputFormat(level Level, message string) string {
	time := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf("[%s] %s %s", levelToString(level), time, message)
	return msg
}

func (l *Logify) log(level Level, message string, calldepth int) error {
	if l.level < level {
		return nil
	}
	logger := log.New(l.out, "", 0)
	msg := outputFormat(level, message)
	return logger.Output(calldepth, msg)
}

func (l *Logify) SetOutput(w io.Writer) {
	l.out = w
}

func (l *Logify) SetLevel(level Level) {
	l.level = level
}

func (l *Logify) Debug(message string) {
	l.log(Debug, message, 0)
}

func (l *Logify) DebugF(format string, i ...interface{}) {
	l.log(Debug, fmt.Sprintf(format, i...), 0)
}

func (l *Logify) Info(message string) {
	l.log(Info, message, 0)
}

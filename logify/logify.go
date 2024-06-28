package logify

import (
	"io"
	"log"
	"os"
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

func (l *Logify) log(level Level, message string, calldepth int) error {
	if l.level < level {
		return nil
	}
	logger := log.New(l.out, "", 0)
	return logger.Output(calldepth, message)
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

func (l *Logify) Info(message string) {
	l.log(Info, message, 0)
}

package logify

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Level uint32

const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
)

func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case Debug:
		return []byte("Debug"), nil
	case Info:
		return []byte("Info"), nil
	case Error:
		return []byte("Error"), nil
	case Fatal:
		return []byte("Fatal"), nil
	}

	return nil, fmt.Errorf("invalid log level: %d", level)
}

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

func formatLogMessage(level Level, message string) string {
	levelText, err := level.MarshalText()
	if err != nil {
		panic(err)
	}

	time := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf("[%s] %s %s", levelText, time, message)
	return msg
}

func (l *Logify) log(level Level, message string, calldepth int) error {
	if l.level > level {
		return nil
	}

	logger := log.New(l.out, "", 0)
	msg := formatLogMessage(level, message)
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

func (l *Logify) InfoF(format string, i ...interface{}) {
	l.log(Info, fmt.Sprintf(format, i...), 0)
}

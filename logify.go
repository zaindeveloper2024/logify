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
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case DebugLevel:
		return []byte("DEBU"), nil
	case InfoLevel:
		return []byte("INFO"), nil
	case WarnLevel:
		return []byte("WARN"), nil
	case ErrorLevel:
		return []byte("ERRO"), nil
	case FatalLevel:
		return []byte("FATA"), nil
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
		level: InfoLevel,
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
	l.log(DebugLevel, message, 0)
}

func (l *Logify) Debugf(format string, i ...interface{}) {
	l.log(DebugLevel, fmt.Sprintf(format, i...), 0)
}

func (l *Logify) Info(message string) {
	l.log(InfoLevel, message, 0)
}

func (l *Logify) Infof(format string, i ...interface{}) {
	l.log(InfoLevel, fmt.Sprintf(format, i...), 0)
}

func (l *Logify) Warn(message string) {
	l.log(WarnLevel, message, 0)
}

func (l *Logify) Warnf(format string, i ...interface{}) {
	l.log(WarnLevel, fmt.Sprintf(format, i...), 0)
}

func (l *Logify) Error(message string) {
	l.log(ErrorLevel, message, 0)
}

func (l *Logify) Errorf(format string, i ...interface{}) {
	l.log(ErrorLevel, fmt.Sprintf(format, i...), 0)
}

func (l *Logify) Fatal(message string) {
	l.log(FatalLevel, message, 0)
}

func (l *Logify) Fatalf(format string, i ...interface{}) {
	l.log(FatalLevel, fmt.Sprintf(format, i...), 0)
}

package logify

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestFormatLogMessage(t *testing.T) {
	level := InfoLevel
	message := "This is a test message"
	expectedTime := time.Now().Format(time.RFC3339)
	expectedPrefix := "[INFO] " + expectedTime

	formattedMessage := formatLogMessage(level, message)

	if !strings.Contains(formattedMessage, expectedPrefix) {
		t.Errorf("expected prefix %s, got %s", expectedPrefix, formattedMessage)
	}

	if !strings.Contains(formattedMessage, message) {
		t.Errorf("expected message %s, got %s", message, formattedMessage)
	}
}

func TestLogify(t *testing.T) {
	buf := new(bytes.Buffer)
	log := New()
	log.SetOutput(buf)
	log.SetLevel(DebugLevel)

	log.Debug("Debug message")
	log.Info("Info message")

	output := buf.String()

	if !strings.Contains(output, "Debug message") {
		t.Errorf("expected Debug message to be logged, got %s", output)
	}

	if !strings.Contains(output, "Info message") {
		t.Errorf("expected Info message to be logged, got %s", output)
	}
}

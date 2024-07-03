package logify

import (
	"strings"
	"testing"
	"time"
)

func TestFormatLogMessage(t *testing.T) {
	level := Info
	message := "This is a test message"
	expectedTime := time.Now().Format(time.RFC3339)
	expectedPrefix := "[Info] " + expectedTime

	formattedMessage := formatLogMessage(level, message)

	if !strings.Contains(formattedMessage, expectedPrefix) {
		t.Errorf("expected prefix %s, got %s", expectedPrefix, formattedMessage)
	}

	if !strings.Contains(formattedMessage, message) {
		t.Errorf("expected message %s, got %s", message, formattedMessage)
	}
}

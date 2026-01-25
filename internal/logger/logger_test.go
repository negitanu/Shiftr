package logger

import (
	"log/slog"
	"testing"
)

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		in   string
		want slog.Level
	}{
		{"DEBUG", slog.LevelDebug},
		{"debug", slog.LevelDebug},
		{"INFO", slog.LevelInfo},
		{"warn", slog.LevelWarn},
		{"WARNING", slog.LevelWarn},
		{"ERROR", slog.LevelError},
		{"", slog.LevelInfo},
		{"unknown", slog.LevelInfo},
	}

	for _, tt := range tests {
		got := ParseLogLevel(tt.in)
		if got != tt.want {
			t.Fatalf("ParseLogLevel(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestGetLogger_Fallback(t *testing.T) {
	orig := logger
	t.Cleanup(func() { logger = orig })

	logger = nil
	got := GetLogger()
	if got == nil {
		t.Fatalf("expected non-nil logger")
	}
	if logger == nil {
		t.Fatalf("expected package logger to be initialized")
	}
}


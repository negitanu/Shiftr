package config

import "testing"

func TestGetDefaultConfig(t *testing.T) {
	cfg := GetDefaultConfig()
	if cfg == nil {
		t.Fatalf("expected non-nil config")
	}
	if cfg.Version == "" {
		t.Fatalf("expected Version to be set")
	}
	if cfg.Settings.LogLevel == "" {
		t.Fatalf("expected Settings.LogLevel to be set")
	}
	// 後方互換: 未設定でも動作する前提
	if cfg.Settings.EnableNotifications != true {
		t.Fatalf("expected EnableNotifications to be true by default")
	}
}


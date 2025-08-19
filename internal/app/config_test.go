package app

import (
	"os"
	"testing"
)

func TestLoadConfig_Defaults(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	cfg, err := LoadConfig(tempDir)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Test default values
	if cfg.App.Name != "wapp-socket" {
		t.Errorf("App.Name = %v, want wapp-socket", cfg.App.Name)
	}

	if cfg.App.LogLevel != "info" {
		t.Errorf("App.LogLevel = %v, want info", cfg.App.LogLevel)
	}

	if cfg.Fakes.Seed != 0 {
		t.Errorf("Fakes.Seed = %v, want 0", cfg.Fakes.Seed)
	}

	if !cfg.Features.FakeMode {
		t.Errorf("Features.FakeMode = %v, want true", cfg.Features.FakeMode)
	}

	if cfg.Features.EventingModel != "bus" {
		t.Errorf("Features.EventingModel = %v, want bus", cfg.Features.EventingModel)
	}
}

func TestLoadConfig_EnvironmentVariables(t *testing.T) {
	// Set environment variables
	os.Setenv("WAPP_APP_NAME", "test-app")
	os.Setenv("WAPP_APP_LOG_LEVEL", "debug")
	os.Setenv("WAPP_FEATURES_FAKE_MODE", "false")
	defer func() {
		os.Unsetenv("WAPP_APP_NAME")
		os.Unsetenv("WAPP_APP_LOG_LEVEL")
		os.Unsetenv("WAPP_FEATURES_FAKE_MODE")
	}()

	tempDir := t.TempDir()
	cfg, err := LoadConfig(tempDir)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if cfg.App.Name != "test-app" {
		t.Errorf("App.Name = %v, want test-app", cfg.App.Name)
	}

	if cfg.App.LogLevel != "debug" {
		t.Errorf("App.LogLevel = %v, want debug", cfg.App.LogLevel)
	}

	if cfg.Features.FakeMode {
		t.Errorf("Features.FakeMode = %v, want false", cfg.Features.FakeMode)
	}
}

func TestLoadConfig_InvalidPath(t *testing.T) {
	// Test with non-existent path - should not fail because file is optional
	cfg, err := LoadConfig("/non/existent/path")
	if err != nil {
		t.Fatalf("LoadConfig should not fail with non-existent path: %v", err)
	}

	// Should still have default values
	if cfg.App.Name != "wapp-socket" {
		t.Errorf("App.Name = %v, want wapp-socket", cfg.App.Name)
	}
}

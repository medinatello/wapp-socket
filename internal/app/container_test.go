package app

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewContainer_Integration(t *testing.T) {
	// Test with default config file setup

	// Create a temporary config directory with a valid config file
	tempDir := t.TempDir()
	configContent := `
app:
  name: "test-app"
  log_level: "info"
fakes:
  seed: 12345
  connect_timeout_chance: 0.1
  connect_fail_chance: 0.05
  ack_latency_ms: 100
  receive_interval_ms: 1000
features:
  fake_mode: true
  eventing_model: "bus"
`

	configPath := filepath.Join(tempDir, "config.yaml")
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Change working directory to temp dir temporarily
	originalWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(originalWd)

	// Create configs directory and copy the config
	configsDir := "configs"
	err = os.Mkdir(configsDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create configs directory: %v", err)
	}

	err = os.WriteFile(filepath.Join(configsDir, "config.yaml"), []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create config in configs dir: %v", err)
	}

	container, err := NewContainer()
	if err != nil {
		t.Fatalf("NewContainer failed: %v", err)
	}

	if container == nil {
		t.Fatal("NewContainer returned nil container")
	}

	// Verify all components are initialized
	if container.Config == nil {
		t.Error("Container.Config is nil")
	}

	if container.Logger == nil {
		t.Error("Container.Logger is nil")
	}

	if container.Telemetry == nil {
		t.Error("Container.Telemetry is nil")
	}

	if container.ConnectUseCase == nil {
		t.Error("Container.ConnectUseCase is nil")
	}

	if container.SendMessageUseCase == nil {
		t.Error("Container.SendMessageUseCase is nil")
	}

	if container.ReceiveUseCase == nil {
		t.Error("Container.ReceiveUseCase is nil")
	}

	if container.GroupsUseCase == nil {
		t.Error("Container.GroupsUseCase is nil")
	}

	// Verify config values
	if container.Config.App.Name != "test-app" {
		t.Errorf("Config.App.Name = %v, want test-app", container.Config.App.Name)
	}

	if container.Config.Fakes.Seed != 12345 {
		t.Errorf("Config.Fakes.Seed = %v, want 12345", container.Config.Fakes.Seed)
	}
}

func TestNewContainer_WithDefaults(t *testing.T) {
	// Test container creation when no config file exists (should use defaults)
	tempDir := t.TempDir()

	// Change working directory to temp dir (no config file)
	originalWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(originalWd)

	// Create empty configs directory
	err := os.Mkdir("configs", 0755)
	if err != nil {
		t.Fatalf("Failed to create configs directory: %v", err)
	}

	container, err := NewContainer()
	if err != nil {
		t.Fatalf("NewContainer failed: %v", err)
	}

	if container == nil {
		t.Fatal("NewContainer returned nil container")
	}

	// Verify default config values
	if container.Config.App.Name != "wapp-socket" {
		t.Errorf("Config.App.Name = %v, want wapp-socket", container.Config.App.Name)
	}

	if container.Config.App.LogLevel != "info" {
		t.Errorf("Config.App.LogLevel = %v, want info", container.Config.App.LogLevel)
	}

	if !container.Config.Features.FakeMode {
		t.Errorf("Config.Features.FakeMode = %v, want true", container.Config.Features.FakeMode)
	}
}

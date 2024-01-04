package config

import (
	"os"
	"testing"
)

func TestConfigProcess(t *testing.T) {
	tempFilePath := "test_config.yaml"
	defer os.Remove(tempFilePath)

	yamlContent := []byte("service:\n  name: custom\n  host: 127.0.0.1\n  port: \"8080\"\n  version: 2.0.0")
	err := os.WriteFile(tempFilePath, yamlContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write YAML file: %v", err)
	}

	configInstance := New(tempFilePath)

	instance, err := configInstance.Process()
	if err != nil {
		t.Errorf("Process returned an error: %v", err)
	}

	instance2, err := configInstance.Process()
	if instance2 != instance {
		t.Error("Expected the same instance for multiple calls to Process()")
	}
}

func TestConfigLoad(t *testing.T) {
	tempFilePath := "test_config.yaml"
	defer os.Remove(tempFilePath)

	yamlContent := []byte("service:\n  name: custom\n  host: 127.0.0.1\n  port: \"8080\"\n  version: 2.0.0")
	err := os.WriteFile(tempFilePath, yamlContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write YAML file: %v", err)
	}

	configInstance := New(tempFilePath)

	err = configInstance.load()
	if err != nil {
		t.Errorf("load returned an error: %v", err)
	}
}

func TestConfigLoadNonExistentFile(t *testing.T) {
	tempFilePath := "nonexistent_config.yaml"
	defer os.Remove(tempFilePath)

	configInstance := New(tempFilePath)

	err := configInstance.load()
	if err == nil {
		t.Error("Expected an error for loading a nonexistent file, but got none")
	}
}

func TestConfigLoadInvalidYAML(t *testing.T) {
	tempFilePath := "invalid_config.yaml"
	defer os.Remove(tempFilePath)

	yamlContent := []byte("invalid_yaml_content")
	err := os.WriteFile(tempFilePath, yamlContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write invalid YAML file: %v", err)
	}

	configInstance := New(tempFilePath)

	err = configInstance.load()
	if err == nil {
		t.Error("Expected an error for loading invalid YAML content, but got none")
	}
}

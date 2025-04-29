// Copyright (c) 2025 Kayvan Sylvan. This project is licensed under the MIT License

package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name: "valid openai config",
			cfg: Config{
				LLM: LLMConfig{
					Provider: "openai",
					APIKey:   "test-key",
					Model:    "gpt-4",
				},
			},
			wantErr: false,
		},
		{
			name: "valid anthropic config",
			cfg: Config{
				LLM: LLMConfig{
					Provider: "anthropic",
					APIKey:   "test-key",
					Model:    "claude-3-opus",
				},
			},
			wantErr: false,
		},
		{
			name: "valid ollama config",
			cfg: Config{
				LLM: LLMConfig{
					Provider: "ollama",
					Endpoint: "http://localhost:11434",
					Model:    "llama3",
				},
			},
			wantErr: false,
		},
		{
			name: "valid lmstudio config",
			cfg: Config{
				LLM: LLMConfig{
					Provider: "lmstudio",
					Endpoint: "http://localhost:1234",
					Model:    "local-model",
				},
			},
			wantErr: false,
		},
		{
			name: "missing api key for cloud provider",
			cfg: Config{
				LLM: LLMConfig{
					Provider: "openai",
					APIKey:   "",
					Model:    "gpt-4",
				},
			},
			wantErr: true,
		},
		{
			name: "missing endpoint for local provider",
			cfg: Config{
				LLM: LLMConfig{
					Provider: "ollama",
					Endpoint: "",
					Model:    "llama3",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid audience",
			cfg: Config{
				LLM: LLMConfig{
					Provider: "openai",
					APIKey:   "test-key",
					Model:    "gpt-4",
				},
				Defaults: DefaultsConfig{
					Audience: "invalid-audience",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variable for testing cloud provider with empty API key
			if tt.name == "missing api key for cloud provider" {
				// First, ensure the environment variable isn't set
				os.Unsetenv("CODEDECODER_LLM_APIKEY")
			}

			err := tt.cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// Test with environment variable set for API key
	t.Run("api key from environment", func(t *testing.T) {
		// Set API key environment variable
		os.Setenv("CODEDECODER_LLM_APIKEY", "env-test-key")
		defer os.Unsetenv("CODEDECODER_LLM_APIKEY")

		cfg := Config{
			LLM: LLMConfig{
				Provider: "openai",
				APIKey:   "", // Empty, but environment variable should cover it
				Model:    "gpt-4",
			},
		}

		err := cfg.Validate()
		if err != nil {
			t.Errorf("Config.Validate() error = %v, expected no error with env var", err)
		}
	})
}

func TestLoadConfig(t *testing.T) {
	// Create a temporary directory for test config files
	tmpDir, err := os.MkdirTemp("", "config-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a test config file
	testConfigContent := `llm:
  provider: openai
  api_key: test-key-from-file
  model: gpt-4
defaults:
  audience: developer
  language: English
  output_dir: ./tutorials
`
	testConfigPath := filepath.Join(tmpDir, "config.yaml")
	if err := os.WriteFile(testConfigPath, []byte(testConfigContent), 0644); err != nil {
		t.Fatalf("Failed to write test config file: %v", err)
	}

	// Test loading from explicit config file
	t.Run("load from explicit config file", func(t *testing.T) {
		cfg, err := LoadConfig(testConfigPath)
		if err != nil {
			t.Fatalf("LoadConfig() error = %v", err)
		}

		if cfg.LLM.Provider != "openai" {
			t.Errorf("Expected provider 'openai', got '%s'", cfg.LLM.Provider)
		}
		if cfg.LLM.APIKey != "test-key-from-file" {
			t.Errorf("Expected API key 'test-key-from-file', got '%s'", cfg.LLM.APIKey)
		}
		if cfg.Defaults.Audience != "developer" {
			t.Errorf("Expected audience 'developer', got '%s'", cfg.Defaults.Audience)
		}
	})

	// Test loading with invalid config file path
	t.Run("load with invalid config path", func(t *testing.T) {
		_, err := LoadConfig(filepath.Join(tmpDir, "nonexistent.yaml"))
		if err == nil {
			t.Error("LoadConfig() expected error with nonexistent file")
		}
	})

	// Test environment variable override
	t.Run("environment variable override", func(t *testing.T) {
		// Set environment variable
		os.Setenv("CODEDECODER_LLM_PROVIDER", "anthropic")
		defer os.Unsetenv("CODEDECODER_LLM_PROVIDER")

		cfg, err := LoadConfig(testConfigPath)
		if err != nil {
			t.Fatalf("LoadConfig() error = %v", err)
		}

		// Check if environment variable overrides config file
		if cfg.LLM.Provider != "anthropic" {
			t.Errorf("Expected provider 'anthropic' from env var, got '%s'", cfg.LLM.Provider)
		}
	})
}

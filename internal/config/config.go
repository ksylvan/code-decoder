// Copyright (c) 2025 Kayvan Sylvan. This project is licensed under the MIT License

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	LLM      LLMConfig      `mapstructure:"llm"`
	Defaults DefaultsConfig `mapstructure:"defaults"`
	GitHub   GitHubConfig   `mapstructure:"github"`
}

// LLMConfig holds configuration for the LLM provider
type LLMConfig struct {
	Provider string `mapstructure:"provider"` // e.g., "openai", "anthropic", "ollama", "lmstudio"
	APIKey   string `mapstructure:"api_key"`  // API key for cloud providers
	Model    string `mapstructure:"model"`    // Specific model to use (e.g., "gpt-4", "claude-3-opus")
	Endpoint string `mapstructure:"endpoint"` // Endpoint URL for local providers (Ollama, LM Studio)
}

// DefaultsConfig holds default settings for operations
type DefaultsConfig struct {
	OutputDir string   `mapstructure:"output_dir"` // Default directory for generated tutorials
	Language  string   `mapstructure:"language"`   // Default tutorial language
	Audience  string   `mapstructure:"audience"`   // Default target audience
	Include   []string `mapstructure:"include"`    // Default include patterns
	Exclude   []string `mapstructure:"exclude"`    // Default exclude patterns
	MaxSize   int64    `mapstructure:"max_size"`   // Default max file size
}

// GitHubConfig holds configuration related to GitHub access
type GitHubConfig struct {
	Token string `mapstructure:"token"` // GitHub personal access token for private repos
}

// LoadConfig reads configuration from file, environment variables, and flags.
// Precedence: Flags > Env > Config File (current dir) > Config File (home dir)
func LoadConfig(cfgFile string) (*Config, error) {
	v := viper.New()

	// 1. Set defaults (optional, if you have hardcoded defaults)
	// v.SetDefault("defaults.output_dir", "./tutorials")
	// v.SetDefault("llm.provider", "openai")

	// 2. Set config file paths
	if cfgFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory and current directory.
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get user home directory: %w", err)
		}
		v.AddConfigPath(filepath.Join(home, ".config", "code-decoder")) // ~/.config/code-decoder/
		v.AddConfigPath(".")                                            // Current directory
		v.SetConfigName("config")                                       // Name of config file (without extension)
		v.SetConfigType("yaml")                                         // REQUIRED if the config file does not have the extension in the name
	}

	// 3. Set environment variable handling
	v.SetEnvPrefix("CODEDECODER") // e.g., CODEDECODER_LLM_PROVIDER
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv() // Read in environment variables that match

	// 4. Read the configuration file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if not explicitly specified via flag
			if cfgFile != "" {
				return nil, fmt.Errorf("config file specified but not found: %s", cfgFile)
			}
			fmt.Println("Config file not found, using defaults and environment variables.")
		} else {
			// Config file was found but another error was produced
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	} else {
		fmt.Println("Using config file:", v.ConfigFileUsed())
	}

	// 5. Unmarshal the config into the struct
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// 6. Bind flags (This should happen in the cmd package where flags are defined)
	// Example: v.BindPFlag("llm.provider", rootCmd.PersistentFlags().Lookup("provider"))
	// Since LoadConfig is called before flags are parsed by Cobra initially,
	// flag values need to be manually checked or Viper needs to be updated after flag parsing.
	// Cobra's OnInitialize handles this timing. Viper instance used in initConfig
	// will have access to parsed flags if BindPFlag was called correctly in cmd's init().
	// For simplicity here, we assume Viper handles precedence correctly if flags are bound.

	// 7. Validate the configuration
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &cfg, nil
}

// Validate checks if the loaded configuration is valid.
func (c *Config) Validate() error {
	// Basic validation example
	if c.LLM.Provider == "" {
		// Depending on commands, this might be acceptable, or it might be an error.
		// For now, just a warning was printed in root.go's initConfig.
		// If required globally, return an error here.
		// return errors.New("llm.provider is required")
	}

	// Add more validation rules as needed
	// e.g., check if API key is present for cloud providers
	isCloudProvider := c.LLM.Provider == "openai" || c.LLM.Provider == "anthropic"
	if isCloudProvider && c.LLM.APIKey == "" {
		// Check environment variable as a fallback before erroring
		envVarName := "CODEDECODER_LLM_APIKEY" // Or specific ones like CODEDECODER_OPENAI_API_KEY
		if os.Getenv(envVarName) == "" {
			return fmt.Errorf("llm.api_key is required for provider '%s' and %s env var is not set", c.LLM.Provider, envVarName)
		}
		// Optionally load from env var directly here if Viper didn't pick it up
		// c.LLM.APIKey = os.Getenv(envVarName)
	}

	isLocalProvider := c.LLM.Provider == "ollama" || c.LLM.Provider == "lmstudio"
	if isLocalProvider && c.LLM.Endpoint == "" {
		return fmt.Errorf("llm.endpoint is required for local provider '%s'", c.LLM.Provider)
	}

	// Validate audience values if necessary
	validAudiences := map[string]bool{"beginner": true, "developer": true, "contributor": true}
	if c.Defaults.Audience != "" && !validAudiences[c.Defaults.Audience] {
		return fmt.Errorf("invalid default audience: '%s'. Must be one of beginner, developer, contributor", c.Defaults.Audience)
	}

	return nil
}

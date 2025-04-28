// Copyright (c) 2025 Kayvan Sylvan. This project is licensed under the MIT License

package cmd

import (
	"fmt"
	"os"

	"github.com/ksylvan/code-decoder/internal/config"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	cfg     *config.Config
	// Root command flags
	versionFlag bool
	// App version set by main
	appVersion string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "code-decoder",
	Short: "AI-Powered Codebase Tutorial Generator",
	Long: `Code-Decoder transforms complex codebases into audience-targeted tutorials using AI.
It analyzes GitHub repositories or local directories, identifies core abstractions,
and generates comprehensive, visualized documentation.`,
	// If the version flag is set, print version and exit.
	// PersistentPreRun removed as version check is now handled in main.go
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	if versionFlag {
	// 		fmt.Printf("code-decoder version %s\n", appVersion)
	// 		os.Exit(0)
	// 	}
	// },
	// Run: func(cmd *cobra.Command, args []string) { }, // Keep commented out unless root command needs direct action
}

// SetVersion allows main to set the version string
func SetVersion(version string) {
	appVersion = version
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent flags (global for application)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/code-decoder/config.yaml or ./config.yaml)")
	// Version flag moved here to be persistent
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "V", false, "Print version information and exit")

	// Local flags (only run when root command is called directly without subcommands)
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.Flags().BoolVarP(&versionFlag, "version", "V", false, "Print version information and exit") // Removed from here
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var err error
	cfg, err = config.LoadConfig(cfgFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v", err)
		// Decide if you want to exit here or let commands handle missing config
		// For now, we'll print the error and continue, commands might not need full config
	}

	// You can add further config validation here if needed
	// For example, check if essential fields are set
	if cfg != nil {
		// Example validation (can be expanded in config.Validate())
		if cfg.LLM.Provider == "" {
			fmt.Fprintln(os.Stderr, "Warning: LLM provider is not set in configuration.")
		}
	} else {
		fmt.Fprintln(os.Stderr, "Warning: Configuration could not be loaded.")
		// Initialize a default config to avoid nil pointer errors later
		cfg = &config.Config{} // Or load defaults
	}
}

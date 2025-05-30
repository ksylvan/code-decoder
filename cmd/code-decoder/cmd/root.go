// Copyright (c) 2025 Kayvan Sylvan. This project is licensed under the MIT License

package cmd

import (
	"fmt"
	"os"

	"github.com/ksylvan/code-decoder/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "V", false, "Print version information and exit")

	// Add the completion command
	rootCmd.AddCommand(completionCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configLoaded := false // Flag to track if any config file was loaded

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err == nil {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
			configLoaded = true
		} else {
			// If the specified config file has an error (e.g., not found, permission denied)
			fmt.Fprintf(os.Stderr, "Error reading specified config file (%s): %s\n", cfgFile, err)
			os.Exit(1) // Exit if the explicitly provided config file fails
		}
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err) // Should not happen normally

		// Search config in home directory and current directory
		viper.AddConfigPath(home + "/.config/code-decoder")
		viper.AddConfigPath(".") // Also look in the current directory
		viper.SetConfigName("config")
		viper.SetConfigType("yaml") // Explicitly set config type

		// Attempt to read the config file from default locations
		if err := viper.ReadInConfig(); err == nil {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
			configLoaded = true
		} else {
			// Only treat ConfigFileNotFoundError as non-fatal for default locations
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				// Config file was found but another error was produced
				fmt.Fprintf(os.Stderr, "Error reading config file (%s): %s\n", viper.ConfigFileUsed(), err)
			}
		}
	}

	// Read environment variables AFTER attempting to load config files
	// Environment variables can override config file settings or provide defaults
	viper.AutomaticEnv() // read in environment variables that match

	// Check if a config file was loaded. If not, print message and exit.
	if !configLoaded && cfgFile == "" { // Only exit if no default config found AND no --config flag used
		fmt.Fprintln(os.Stderr, "Error: Configuration file not found.")
		fmt.Fprintln(os.Stderr, "Please create a config.yaml in the current directory (./config.yaml)")
		fmt.Fprintln(os.Stderr, "or in your home config directory (~/.config/code-decoder/config.yaml).")
		fmt.Fprintln(os.Stderr, "An example configuration can be found at 'example/config.yaml'.")
		fmt.Fprintln(os.Stderr, "Alternatively, specify a config file using the --config flag.")
		os.Exit(1)
	}
}

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate the autocompletion script for the specified shell",
	Long: `To load completions:

Bash:

  $ source <(code-decoder completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ code-decoder completion bash > /etc/bash_completion.d/code-decoder
  # macOS:
  $ code-decoder completion bash > /usr/local/etc/bash_completion.d/code-decoder

Zsh:

  # If shell completion is not already enabled in your environment, you will need
  # to enable it. You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ code-decoder completion zsh > "${fpath[1]}/_code-decoder"

  # You will need to start a new shell for this setup to take effect.

Fish:

  $ code-decoder completion fish | source

  # To load completions for each session, execute once:
  $ code-decoder completion fish > ~/.config/fish/completions/code-decoder.fish

PowerShell:

  PS> code-decoder completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> code-decoder completion powershell > code-decoder.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		shells := map[string]func(*cobra.Command, os.FileInfo){
			"bash":       func(c *cobra.Command, f os.FileInfo) { c.GenBashCompletion(os.Stdout) },
			"zsh":        func(c *cobra.Command, f os.FileInfo) { c.GenZshCompletion(os.Stdout) },
			"fish":       func(c *cobra.Command, f os.FileInfo) { c.GenFishCompletion(os.Stdout, true) },
			"powershell": func(c *cobra.Command, f os.FileInfo) { c.GenPowerShellCompletionWithDesc(os.Stdout) },
		}
		if gen := shells[args[0]]; gen != nil {
			gen(cmd.Root(), nil)
		} else {
			fmt.Fprintf(os.Stderr, "Unknown shell: %s\n", args[0])
			os.Exit(1)
		}
	},
}

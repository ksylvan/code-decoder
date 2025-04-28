// Copyright (c) 2025 Kayvan Sylvan. This project is licensed under the MIT License

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testLlmCmd represents the test-llm command
var testLlmCmd = &cobra.Command{
	Use:   "test-llm",
	Short: "Test the connection to the configured LLM provider",
	Long: `Verifies that the application can successfully connect to the
Large Language Model (LLM) provider specified in the configuration
(or overridden via flags) and perform a basic interaction.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test-llm called")
		providerOverride, _ := cmd.Flags().GetString("provider")
		fmt.Printf("Testing LLM connection (Provider from config: %s, Override: %s)", cfg.LLM.Provider, providerOverride)

		// TODO: Implement LLM connection test logic
		// 1. Determine the provider to use (config or override)
		// 2. Get provider configuration (API key, endpoint, model)
		// 3. Initialize the LLM client/provider
		// 4. Call the provider's TestConnection method or perform a simple API call
		// 5. Report success or failure
	},
}

func init() {
	rootCmd.AddCommand(testLlmCmd)

	// Flags for test-llm command
	testLlmCmd.Flags().String("provider", "", "Override the LLM provider specified in the config for this test")

	// TODO: Bind flags if needed, though provider override is handled directly in Run for now
	// viper.BindPFlag("testllm.provider", testLlmCmd.Flags().Lookup("provider"))
}

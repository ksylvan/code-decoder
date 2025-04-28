// Copyright (c) 2025 Kayvan Sylvan. This project is licensed under the MIT License

package cmd

import (
	"fmt"
	"os" // Added for error handling in completion registration

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate tutorials from a codebase or saved analysis",
	Long: `Creates audience-targeted tutorials based on either a direct codebase analysis
or a previously saved analysis file. Outputs can be customized by audience,
language, and format.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		// TODO: Implement generation logic here
		// 1. Determine source: load analysis or analyze dir/repo
		// 2. Get generation options (audience, language, format, output dir)
		// 3. Generate content using LLM and analysis data
		// 4. Render content using templates (Markdown/HTML)
		// 5. Save output files
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Flags for generate command
	generateCmd.Flags().String("load-analysis", "", "Path to a saved analysis file to use for generation")
	generateCmd.Flags().String("dir", "", "Path to the local directory to analyze and generate from")
	generateCmd.Flags().String("repo", "", "URL of the GitHub repository to analyze and generate from")
	generateCmd.Flags().String("audience", "developer", "Target audience for the tutorial (beginner, developer, contributor)")
	generateCmd.Flags().String("language", "English", "Language for the generated tutorial")
	generateCmd.Flags().String("output", "./tutorials", "Directory to save generated tutorials")
	generateCmd.Flags().String("format", "markdown", "Output format (markdown, html)")
	generateCmd.Flags().String("save-analysis", "", "File path to save analysis results if analyzing a codebase directly")
	generateCmd.Flags().String("provider", "", "Override the LLM provider specified in the config")
	generateCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")

	// Register custom completion for the --audience flag
	err := generateCmd.RegisterFlagCompletionFunc("audience", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"beginner", "developer", "contributor"}, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		// Handle error, e.g., log it or exit. Exiting is simple for init phase.
		fmt.Fprintf(os.Stderr, "Error registering completion function for --audience: %v\n", err)
		os.Exit(1)
	}

	// Ensure either load-analysis or one of (dir, repo) is provided
	generateCmd.MarkFlagsMutuallyExclusive("load-analysis", "dir")
	generateCmd.MarkFlagsMutuallyExclusive("load-analysis", "repo")
	// Note: dir and repo are already mutually exclusive via analyzeCmd logic if we reuse it,
	// but explicit here is fine too. If generate directly analyzes, it needs this.
	generateCmd.MarkFlagsMutuallyExclusive("dir", "repo")

	// TODO: Add logic to ensure that if --load-analysis is not used, then either --dir or --repo MUST be provided.
	// This might require validation within the Run function or using PersistentPreRunE.
}

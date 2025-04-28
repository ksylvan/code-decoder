// Copyright (c) 2025 Kayvan Sylvan. This project is licensed under the MIT License

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze a codebase and save the analysis",
	Long: `Processes a codebase from a local directory or GitHub repository,
extracts structural information and high-level knowledge,
and saves the analysis to a specified file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("analyze called")
		// TODO: Implement analysis logic here
		// 1. Get source (dir or repo)
		// 2. Validate source
		// 3. List files based on config (include/exclude/size)
		// 4. Parse files
		// 5. Extract knowledge using LLM
		// 6. Save analysis to file
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	// Flags for analyze command
	analyzeCmd.Flags().String("dir", "", "Path to the local directory to analyze")
	analyzeCmd.Flags().String("repo", "", "URL of the GitHub repository to analyze")
	analyzeCmd.Flags().String("save-analysis", "", "File path to save the analysis results (required)")
	analyzeCmd.Flags().String("name", "", "Custom project name")
	analyzeCmd.Flags().String("token", "", "GitHub token for private repositories")
	analyzeCmd.Flags().StringSlice("include", nil, "File patterns to include (comma-separated or multiple flags)")
	analyzeCmd.Flags().StringSlice("exclude", nil, "File patterns to exclude (comma-separated or multiple flags)")
	analyzeCmd.Flags().Int64("max-size", 0, "Maximum file size in bytes to include")
	analyzeCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")

	// Mark save-analysis as required - Commenting out to test completion behavior
	// analyzeCmd.MarkFlagRequired("save-analysis")

	// Ensure either --dir or --repo is provided, but not both
	analyzeCmd.MarkFlagsMutuallyExclusive("dir", "repo")

	// TODO: Bind flags to Viper configuration for analyze-specific overrides if needed
	// viper.BindPFlag("analyze.output", analyzeCmd.Flags().Lookup("save-analysis"))
}

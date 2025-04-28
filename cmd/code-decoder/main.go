// Copyright (c) 2025 Kayvan Sylvan. This project is licensed under the MIT License

package main

import (
	"fmt"
	"os"

	"github.com/ksylvan/code-decoder/cmd/code-decoder/cmd"
)

// version is set during build time using ldflags
var version = "dev" // Default value

func main() {
	// Check for version flag before doing anything else
	for _, arg := range os.Args[1:] {
		if arg == "-V" || arg == "--version" {
			fmt.Printf("code-decoder version %s\n", version)
			os.Exit(0)
		}
	}

	// Pass the version to the command package
	cmd.SetVersion(version)
	cmd.Execute()
}

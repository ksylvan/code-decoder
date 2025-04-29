// Copyright (c) 2025 Kayvan Sylvan. This project is licensed under the MIT License

package main

import (
	"os"
	"testing"
)

func TestMainVersionFlag(t *testing.T) {
	// Save original arguments and restore after test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Redirect stdout to capture output
	oldStdout := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	// Restore stdout after test
	defer func() {
		w.Close()
		os.Stdout = oldStdout
	}()

	// Mock args with version flag
	os.Args = []string{"code-decoder", "--version"}

	// Execute function that should just exit without error
	// This just tests that code compiles and doesn't panic
	// The actual exit behavior can't be tested directly

	// We don't want to actually run main() as it would call os.Exit()
	// Just verify the version variable exists
	if version == "" {
		t.Error("Version string should not be empty")
	}
}

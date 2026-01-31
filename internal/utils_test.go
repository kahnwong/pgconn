package internal

import (
	"testing"
)

func TestCheckIfBinaryExists(t *testing.T) {
	// Test with a binary that should exist on all systems
	// This test just verifies the function doesn't panic for existing binaries
	// We can't easily test the failure case since it calls os.Exit(1)

	// Just verify the function can be called without panicking
	// Using "sh" which should exist on all Unix-like systems
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("CheckIfBinaryExists panicked: %v", r)
		}
	}()

	CheckIfBinaryExists("sh")
}

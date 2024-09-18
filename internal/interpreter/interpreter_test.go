package interpreter_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/AlexanderStocks/GoGo/internal/interpreter"
)

func TestRunFile(t *testing.T) {
	testFiles := []string{
		"../../testdata/test_println.go",
		"../../testdata/test_assignment.go",
		"../../testdata/test_if.go",
		// "../../testdata/test_for.go",
	}

	for _, file := range testFiles {
		t.Run(filepath.Base(file), func(t *testing.T) {
			err := interpreter.RunFile(file)
			if err != nil {
				t.Errorf("interpreter error for %s: %v", file, err)
			}
		})
	}
}

func writeTempFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func removeTempFile(filename string) {
	_ = os.Remove(filename)
}

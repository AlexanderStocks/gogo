package interpreter_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/AlexanderStocks/GoGo/internal/interpreter"
)

func TestRunFile(t *testing.T) {
	testFiles := []string{
		"../../testdata/test1.go",
		"../../testdata/test2.go",
		// "../../testdata/test3.go",
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

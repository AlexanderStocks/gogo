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
		// "../../testdata/test2.go",
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

func TestRunFile_MissingMain(t *testing.T) {
	src := `
        package main
        func notMain() {}
    `
	file := "../../testdata/missing_main.go"
	err := writeTempFile(file, src)
	if err != nil {
		t.Fatalf("could not write temp file: %v", err)
	}
	defer removeTempFile(file)

	err = interpreter.RunFile(file)
	if err == nil {
		t.Error("expected error for missing main function, got nil")
	}
}

func writeTempFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func removeTempFile(filename string) {
	_ = os.Remove(filename)
}

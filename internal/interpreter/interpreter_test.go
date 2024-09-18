package interpreter_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/AlexanderStocks/GoGo/internal/interpreter"
)

func captureOutput(f func()) string {
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = origStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestRunFile(t *testing.T) {
	testFiles := []struct {
		name     string
		filename string
		expected string
	}{
		{"Println Test", "../../testdata/test_println.go", "Test 1: Hello, World!\n"},
		{"Assignment Test", "../../testdata/test_assignment.go", "Test 2: x * y = 6\n"},
		{"If Statement Test", "../../testdata/test_if.go", "Test 3: i = 0\nTest 3: i = 1\nTest 3: i = 2\n"},
		{"For Loop Test", "../../testdata/test_for.go", "Test 3: i = 0\nTest 3: i = 1\nTest 3: i = 2\n"},
	}

	for _, tt := range testFiles {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			output := captureOutput(func() {
				err := interpreter.RunFile(tt.filename)
				if err != nil {
					t.Fatalf("interpreter error: %v", err)
				}
			})

			output = strings.ReplaceAll(output, "\r\n", "\n") // Handle Windows line endings

			if output != tt.expected {
				t.Errorf("expected output:\n%v\ngot:\n%v", tt.expected, output)
			}
		})
	}
}

package evaluatortest

import (
	"testing"
)

func TestEvalStatements(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{"Variable Declaration", "var result = 5", int64(5)},
		{"Variable Assignment", "var x = 5\nx = 10\nvar result = x", int64(10)},
		{"Multiple Assignment", "var x, y = 1, 2\nvar result = x + y", int64(3)},
		{"If Statement True", "var result = 0\nif true { result = 1 }", int64(1)},
		{"If Statement False", "var result = 0\nif false { result = 1 }", int64(0)},
		{"If-Else Statement", "var result = 0\nif false { result = 1 } else { result = 2 }", int64(2)},
		{"If-ElseIf-Else Statement", "var result = 0\nif false { result = 1 } else if true { result = 2 } else { result = 3 }", int64(2)},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := testEval(t, tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v (%T), got %v (%T)", tt.expected, tt.expected, result, result)
			}
		})
	}
}

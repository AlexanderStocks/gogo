package evaluatortest

import (
	"testing"
)

func TestEvalAssignment(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{"Simple Assignment", "var x = 5\nx = 10\nvar result = x", int64(10)},
		// {"Multiple Assignment", "var x, y = 1, 2\nx, y = y, x\nvar result = x * 10 + y", int64(21)},
		{"Assignment with Expression", "var x = 5\nx = x + 5\nvar result = x", int64(10)},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := testEval(t, tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

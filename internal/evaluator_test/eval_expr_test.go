package evaluatortest

import (
	"testing"
)

func TestEvalExpressions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{"Integer Addition", "var result = 1 + 2", int64(3)},
		{"Integer Subtraction", "var result = 10 - 5", int64(5)},
		{"Integer Multiplication", "var result = 2 * 3", int64(6)},
		{"Integer Division", "var result = 8 / 2", int64(4)},
		{"Float Addition", "var result = 1.5 + 2.5", 4.0},
		{"Float Multiplication", "var result = 2.0 * 3.5", 7.0},
		{"String Concatenation", `var result = "hello" + " " + "world"`, "hello world"},
		// {"Boolean AND", "var result = true && false", false},
		// {"Boolean OR", "var result = true || false", true},
		// {"Equality Check", "var result = (5 == 5)", true},
		// {"Inequality Check", "var result = (5 != 3)", true},
		// {"Less Than Check", "var result = (3 < 5)", true},
		// {"Greater Than Check", "var result = (5 > 3)", true},
		// {"Unary Minus", "var result = -5", int64(-5)},
		// {"Unary Minus Float", "var result = -3.5", -3.5},
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

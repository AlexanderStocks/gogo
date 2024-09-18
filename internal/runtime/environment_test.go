package runtime

import (
	"testing"
)

func TestEnvironmentGetSet(t *testing.T) {
	env := NewEnvironment(nil)

	env.Set("a", 1)
	env.Set("b", "hello")

	aVal, err := env.Get("a")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if aVal != 1 {
		t.Errorf("expected 1, got %v", aVal)
	}

	bVal, err := env.Get("b")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bVal != "hello" {
		t.Errorf("expected 'hello', got %v", bVal)
	}

	_, err = env.Get("c")
	if err == nil {
		t.Errorf("expected error for undefined variable 'c', got nil")
	}
}

func TestEnvironmentNested(t *testing.T) {
	globalEnv := NewEnvironment(nil)
	globalEnv.Set("a", 1)

	localEnv := NewEnvironment(globalEnv)
	localEnv.Set("b", 2)

	aVal, err := localEnv.Get("a")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if aVal != 1 {
		t.Errorf("expected 1, got %v", aVal)
	}

	bVal, err := localEnv.Get("b")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bVal != 2 {
		t.Errorf("expected 2, got %v", bVal)
	}
}

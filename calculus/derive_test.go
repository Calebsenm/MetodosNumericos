package calculus

import (
	"testing"
)

func TestDerive(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"x^2", "2*x"},
		{"3*x", "3"},
		{"sin(x)", "cos(x)"},
		{"cos(x)", "-sin(x)"},
		{"5", "0"},
	}

	for _, tt := range tests {
		got, err := Derive(tt.input)
		if err != nil {
			t.Errorf("Error al derivar %q: %v", tt.input, err)
		}
		if got != tt.expected {
			t.Errorf("Derive(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
} 
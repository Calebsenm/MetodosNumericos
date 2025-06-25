package calculus

import (
	"testing"
	"math"
)

func TestEvaluateFunction(t *testing.T) {
	tests := []struct {
		expr     string
		x        float64
		expected float64
	}{
		{"x+2", 3, 5},
		{"2*x", 4, 8},
		{"sin(x)", math.Pi / 2, 1},
		{"cos(x)", 0, 1},
		{"x^2", 2, 4},
	}

	for _, tt := range tests {
		got, err := EvaluateFunction(tt.x, tt.expr)
		if err != nil {
			t.Errorf("Error al evaluar %q: %v", tt.expr, err)
		}
		if math.Abs(got-tt.expected) > 1e-6 {
			t.Errorf("EvaluateFunction(%v, %q) = %v, want %v", tt.x, tt.expr, got, tt.expected)
		}
	}
} 
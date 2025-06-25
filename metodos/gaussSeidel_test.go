package metodos

import (
	"testing"
	"math"
)

func TestSolveSystem(t *testing.T) {
	A := [][]float64{
		{4, 1, 2},
		{3, 5, 1},
		{1, 1, 3},
	}
	b := []float64{4, 7, 3}

	iterations, desc := SolveSystem(A, b)

	if len(iterations) == 0 {
		t.Fatal("No se generaron iteraciones")
	}

	if desc == "" {
		t.Error("No se generó la descripción del sistema")
	}

	// Tomamos la última iteración como solución aproximada
	sol := iterations[len(iterations)-1]
	expected := []float64{0.5, 1.0, 0.5}
	for i := range sol {
		if math.Abs(sol[i]-expected[i]) > 1e-2 {
			t.Errorf("La variable %d es incorrecta: got %v, want %v", i, sol[i], expected[i])
		}
	}
} 
package metodos

import (
	"fmt"
	"math"
)

func SolveSystem(A [][]float64, b []float64) (iterations [][]float64, systemDescription string) {
	n := len(b)
	x := make([]float64, n)
	xNew := make([]float64, n)
	iterations = [][]float64{}

	// Pivotación parcial
	pivot(A, b)

	// Generar la descripción del sistema
	systemDescription = "System of equations:\n"
	for i := range A {
		row := ""
		for j := range A[i] {
			row += fmt.Sprintf("%3g*x%d ", A[i][j], j+1)
			if j < len(A[i])-1 {
				row += "+ "
			}
		}
		systemDescription += fmt.Sprintf("[%s] = [%3g]\n", row, b[i])
	}

	// Resolver usando el método Gauss-Seidel con pivotación
	for itCount := 1; itCount <= 1000; itCount++ {
		copy(xNew, x)
		for i := range A {
			// Calculando la suma de las ecuaciones
			s1 := 0.0
			s2 := 0.0
			for j := 0; j < i; j++ {
				s1 += A[i][j] * xNew[j]
			}
			for j := i + 1; j < n; j++ {
				s2 += A[i][j] * x[j]
			}
			// Actualizando el valor de x[i]
			xNew[i] = (b[i] - s1 - s2) / A[i][i]
		}
		iterations = append(iterations, append([]float64(nil), xNew...)) // Guardar copia de la iteración

		// Verificar si las soluciones actuales son suficientemente cercanas
		if allClose(x, xNew, 1e-8) {
			break
		}
		copy(x, xNew)
	}


	return iterations, systemDescription
}

// Pivotación parcial: Reorganiza las filas de la matriz para que el valor más grande esté en la diagonal principal
func pivot(A [][]float64, b []float64) {
	n := len(A)
	for i := 0; i < n; i++ {
		maxRow := i
		for j := i + 1; j < n; j++ {
			if math.Abs(A[j][i]) > math.Abs(A[maxRow][i]) {
				maxRow = j
			}
		}

		if maxRow != i {
			A[i], A[maxRow] = A[maxRow], A[i]
			b[i], b[maxRow] = b[maxRow], b[i]
		}
	}
}

func allClose(a, b []float64, tol float64) bool {
	for i := range a {
		if math.Abs(a[i]-b[i]) > tol {
			return false
		}
	}
	return true
}
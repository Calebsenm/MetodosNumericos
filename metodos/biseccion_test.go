package metodos

import "testing"

func TestBi(t *testing.T) {

	bi := Biseccion{NumberOne: 0, NumberTwo: 1, MaxIter: 30, Function: "2*x^2 + 2*x -2"}
	data, err := bi.Calcular()

	if err != nil {
		t.Error(err)
	}

	// Verifica que la última aproximación esté cerca de la raíz real
	var bestXi, bestFxi float64
	first := true
	for _, valores := range data {
		fxi := valores[5]
		xi := valores[4]
		if first || abs(fxi) < abs(bestFxi) {
			bestFxi = fxi
			bestXi = xi
			first = false
		}
	}
	if bestXi < 0.617 || bestXi > 0.619 {
		t.Errorf("La raíz aproximada está fuera del rango esperado: got %v", bestXi)
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

package metodos

import (
	"fmt"
	"testing"
)

func TestRe(t *testing.T) {
	regla := ReglaFalsa{
		NumberOne: 0,
		NumberTwo: 1,
		MaxIter:   30,
		Function:  "2*x^2 +2*x -2 ",
	}

	data, err := regla.Calcular()
	if err != nil {
		t.Fatalf("Error al calcular: %v", err)
	}

	// Asegúrate de que el mapa data no esté vacío antes de acceder a los valores
	if len(data) == 0 {
		t.Fatal("El mapa data está vacío")
	}

	// Verificar que todas las claves tengan al menos un elemento
	for k, v := range data {
		fmt.Println(k, v)
	}

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

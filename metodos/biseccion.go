package metodos

import (
	"errors"
	"fmt"
	"metodosNumericos/calculus"
)

type Biseccion struct {
	NumberOne float64
	NumberTwo float64
	MaxIter   int
	Function  string
}

func (mt *Biseccion) Calcular() (map[int][]float64, error) {

	a := mt.NumberOne
	b := mt.NumberTwo

	var xi float64
	data := make(map[int][]float64)

	for i := 0; i < mt.MaxIter; i++ {

		fa, err := calculus.EvaluateFunction(a, mt.Function)
		if err != nil {
			fmt.Println(err)
		}

		fb, err := calculus.EvaluateFunction(b, mt.Function)

		if fa == 0 {
			valores := []float64{a, b, fa, fb, a, fa}
			data[i] = valores
			break
		}
		if fb == 0 {
			valores := []float64{a, b, fa, fb, b, fb}
			data[i] = valores
			break
		}
		if fa*fb > 0 {
			return nil, errors.New("el intervalo no encierra una raíz, elige otro intervalo")
		}

		xi = (a + b) / 2

		fxi, err := calculus.EvaluateFunction(xi, mt.Function)
		if err != nil {
			fmt.Println("error", err)
		}

		valores := []float64{
			a, b, fa, fb, xi, fxi,
		}

		data[i] = valores

		if fa*fxi < 0 {
			b = xi
		} else {
			a = xi
		}
	}

	return data, nil
}

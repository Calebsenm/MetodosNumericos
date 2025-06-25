package metodos

import "testing"

func TestNew(t *testing.T) {

	bi := NewthonRp{Number: 1.0, MaxIter: 10, Function: "2*x^2 + 2*x -2"}
	result, err := bi.Calular()

	if err != nil {
		t.Error(err)
	}

	bestXi := result[len(result)-1]
	if bestXi < 0.617 || bestXi > 0.619 {
		t.Errorf("La raíz aproximada está fuera del rango esperado: got %v", bestXi)
	}
}

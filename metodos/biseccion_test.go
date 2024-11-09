package metodos 

import "testing"

func TestBi(t *testing.T){

    bi := Biseccion{NumberOne: 0 , NumberTwo: 1 ,MaxIter: 10 , Function: "2x^2 + 2x -2"}
    _ , err := bi.Calcular()
    
    if err != nil {
        t.Error(err)
    }

    /*
    if result[0] == 0.5 {
        t.Error("incorrect result")
    }*/
}

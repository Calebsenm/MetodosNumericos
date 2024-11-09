package metodos 

import "testing"

func TestRe(t *testing.T){

    bi := ReglaFalsa{NumberOne: 1 , NumberTwo: 2 ,MaxIter: 10 , Function: "x^3 -x -2"}
    result , err := bi.Calcular()
    
    if err != nil {
        t.Error(err)
    }

    if result[0] != 1.333 {
        t.Error("incorrect result", result  )
    }
}

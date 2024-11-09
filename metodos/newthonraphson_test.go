package metodos 

import "testing"

func TestNew(t *testing.T){

    bi := NewthonRp{Number: 1 ,MaxIter: 10 , Function: "2*x + 2x +2"}
    result , err := bi.Calular()
    
    if err != nil {
        t.Error(err)
    }

    if result[0] == 0.5 {
        t.Error("incorrect result")
    }
}

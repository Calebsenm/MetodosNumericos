package metodos 

import (
    "metodosNumericos/calculus"
)

type NewthonRp struct {
    Number   float64
    Function string
    MaxIter  int 
}

func (mt *NewthonRp) Calular() ([]float64 , error  ){ 
    
    dx , _ := calculus.Derive(mt.Function) 
    var xi float64  = mt.Number  
    var interations []float64

    for i := 0; i < mt.MaxIter ; i++ {
        fa , _ := evaluate( mt.Number , mt.Function) 
        dxfa , _  := evaluate(mt.Number , dx)
    
        xi = xi - ( fa / dxfa)
        interations = append(interations , xi)
        
    }
    return interations , nil  
}



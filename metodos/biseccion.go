package metodos

import (
    "errors"
    "metodosNumericos/calculus"
)

type Biseccion struct {
    NumberOne float64 
    NumberTwo float64
    MaxIter   int 
    Function  string 
}

func (mt *Biseccion) Calcular() ([]float64, error ) {
    
    a := mt.NumberOne 
    b := mt.NumberTwo

    var xi float64 
    var iterations []float64

    for i := 0 ; i < mt.MaxIter ; i++{
                
        fa , _ := calculus.EvaluateFunction( a , mt.Function );
        fb , _ := calculus.EvaluateFunction( b , mt.Function );


        if fa * fb >= 0 {
            return nil   , errors.New("el intervalo no encierra una raíz, elige otro intervalo"); 
        }

        xi =  fa + ( fb - fa)  / 2
        iterations = append( iterations , xi);
            
	    if fa * xi < 0 {
		    b = xi  
        
		} else {
			a = xi 
        
		}
    }
   
    return iterations  , nil  
}



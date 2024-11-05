package metodos

import (
    "errors"
    "metodosNumericos/calculus"
)

type ReglaFalsa struct {
    NumberOne float64 
    NumberTwo float64
    MaxIter   int 
    Function  string 
}

func (mt *ReglaFalsa ) Calcular() ([]float64 , error){
 
    a := mt.NumberOne 
    b := mt.NumberTwo

    var xi float64 
    var iterations []float64

    for i := 0 ; i < mt.MaxIter ; i++{
             
        fa , _  := calculus.EvaluateFunction( a , mt.Function )
        fb , _  := calculus.EvaluateFunction( b , mt.Function )
        
        if fa * fb > 0 {
            return nil   , errors.New("el intervalo no encierra una raíz, elige otro intervalo"); 
        }

        xi = (a * fb  - b * fa  ) / ( fb - fa)
  
        iterations = append( iterations , xi);

        fxi , _ := calculus.EvaluateFunction(xi , mt.Function )

	    if fa * fxi  < 0 {
		    b = xi  
        
		} else {
			a = xi  
		}
    }
   
    return iterations  , nil    
}

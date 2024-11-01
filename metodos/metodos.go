package metodos

import (
   "errors"
 //  "fmt" 
)

type Method struct {
    NumberOne float64 
    NumberTwo float64
    MaxIter   int 
    Function  string 
}

func (mt *Method ) Biseccion() ([]float64, error ) {
    
    a := mt.NumberOne 
    b := mt.NumberTwo

    var xi float64 
    var iterations []float64

    for i := 0 ; i < mt.MaxIter ; i++{
            
    
        fa , _  := evaluate( a , mt.Function )
        fb , _  := evaluate( b , mt.Function )
    
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

func (mt *Method) ReglaFalsa() ([]float64 , error){
 
    a := mt.NumberOne 
    b := mt.NumberTwo

    var xi float64 
    var iterations []float64

    for i := 0 ; i < mt.MaxIter ; i++{
             
        fa , _  := evaluate( a , mt.Function )
        fb , _  := evaluate( b , mt.Function )
        

        if fa * fb > 0 {
            return nil   , errors.New("el intervalo no encierra una raíz, elige otro intervalo"); 
        }

        xi = (a * fb  - b * fa  ) / ( fb - fa)
  
        iterations = append( iterations , xi);

        fxi , _ := evaluate(xi , mt.Function )

	    if fa * fxi  < 0 {
		    b = xi  
        
		} else {
			a = xi  
		}
    }
   
    return iterations  , nil    
}

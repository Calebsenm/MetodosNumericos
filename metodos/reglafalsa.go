package metodos

import (
    "errors"
    "metodosNumericos/calculus"
    "fmt"     
)

type ReglaFalsa struct {
    NumberOne float64 
    NumberTwo float64
    MaxIter   int 
    Function  string 
}

func (mt *ReglaFalsa ) Calcular() (map[int][]float64 , error){
 
    a := mt.NumberOne 
    b := mt.NumberTwo 
    var xi float64 
    data := make(map[int][]float64) 


    for i := 1 ; i < mt.MaxIter ; i++{
             
        fa , err  := calculus.EvaluateFunction( a , mt.Function )
        if err != nil {
            fmt.Println(err)
        }

        fb , err  := calculus.EvaluateFunction( b , mt.Function )
        if err != nil {
            fmt.Println(err)
        } 

        if fa * fb > 0 {
            return nil   , errors.New("el intervalo no encierra una raíz, elige otro intervalo"); 
        }

        xia := a * fb 
        xib := b * fa 
        xif := fb - fa

        xi =  (xia  -  xib) / xif  
        fxi  , err  := calculus.EvaluateFunction(xi , mt.Function)
        
        if err != nil {
            fmt.Println(err)
        }

        valores := []float64{
            a , b , fa , fb , xi, fxi, 
        }  
        
        data[i] =  valores 


	    if fa * fxi  < 0 {
		    b = xi  
        
		} else {
			a = xi  
		}
    }
   
    return data , nil    
}

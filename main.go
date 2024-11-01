
package main 

import (
    "fmt"
    "metodosNumericos/metodos"
)


func main(){
    fmt.Println("---------------- Metodos --------------------");

    var function string = "2* x^2 +2 * x -2 "
    //var function  string = "x^3 -x - 2"
    

    //numbers := metodos.Method{ NumberOne: 1, NumberTwo: 2 ,MaxIter: 10  , Function: function}
    newthon := metodos.NewthonRp{ Number: 0.75, Function: function, MaxIter: 10 }


    //data , _ :=  numbers.Biseccion();
   // data  , _ := numbers.ReglaFalsa();
    data , _:=  newthon.NewThonRaphson();

    for i, value := range data {
        fmt.Printf("Iteración %d: %.4f\n", i+1, value)
    }

        
}

package main 

import (
    "fmt"
    "metodosNumericos/metodos"
)

func main(){
    fmt.Println("---------------- Metodos --------------------");
    
    var opciones map[string]string = map[string] string {} 
    opciones["1"] = "Biseccion";
    opciones["2"] = "Regla Falsa";
    opciones["3"] = "Newthon Raphon";

    for a , b := range opciones {
        fmt.Println(a , b);
    } 
    
    var option int    
    fmt.Print("Selecciona: ");
    fmt.Scan(&option)

    if option == 1 {
        var num1 float64 
        fmt.Print("ingresa el num1: ");
        fmt.Scan(&num1)

        fmt.Print("ingresa el num2: ");
        var num2 float64 
        fmt.Scan(&num2)

        fmt.Print("ingresa el numero interaciones: ")
        var numIter int 
        fmt.Scan(&numIter)

        fmt.Print("ingresa la funcion: ");
        var function string 
        fmt.Scan(&numIter)
        
        bs := metodos.Biseccion{NumberOne: num1 , NumberTwo: num2 , MaxIter: numIter , Function: function} 
        data , err := bs.Calcular()
        
        imprimirData(data);

        if err != nil {
            fmt.Println(err)
        }


    }   else if option == 2{
        
        var num1 float64 
        fmt.Print("ingresa el num1: ");
        fmt.Scan(&num1)

        fmt.Print("ingresa el num2: ");
        var num2 float64 
        fmt.Scan(&num2)

        fmt.Print("ingresa el numero interaciones: ")
        var numIter int 
        fmt.Scan(&numIter)

        fmt.Print("ingresa la funcion: ");
        var function string 
        fmt.Scan(&numIter)
        
        bs := metodos.ReglaFalsa{NumberOne: num1 , NumberTwo: num2 , MaxIter: numIter , Function: function} 
        data , err := bs.Calcular()
        
        imprimirData(data);

        if err != nil {
            fmt.Println(err)
        }


    }   else if option == 3{

        var num1 float64; 
        fmt.Print("Ingresa el numero: ");
        fmt.Scan(&num1);

        fmt.Print("ingresa el numero interaciones: ");
        var numIter int;
        fmt.Scan(&numIter)

        fmt.Print("ingresa la funcion: ");
        var function string 
        fmt.Scan(&numIter)
        
        bs := metodos.NewthonRp{Number: num1 , MaxIter: numIter , Function: function} 
        data , err := bs.Calular()
        
        imprimirData(data);

        if err != nil {
            fmt.Println(err)
        }

    }   else {
            fmt.Println("Opcion no existe")
    }

}

func imprimirData(data []float64){
        
    for i, value := range data {
        fmt.Printf("Iteración %d: %.4f\n", i+1, value)
    }

}

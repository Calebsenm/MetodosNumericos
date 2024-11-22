package main 

import (
    "fmt"
    "metodosNumericos/metodos"
	"bufio"
	"os"
	"strconv"
	"strings"
)


func main(){
    fmt.Println("---------------- Metodos --------------------");
    
    var opciones map[string]string = map[string] string {} 
    opciones["1"] = "Biseccion";
    opciones["2"] = "Regla Falsa";
    opciones["3"] = "Newthon Raphon";
    opciones["4"] = "Gauss Sediel"

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
        _ , err := bs.Calcular()
        
        //imprimirData(data);

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
        _ , err := bs.Calular()
        
        if err != nil {
            fmt.Println(err)
        }

    }   else if option == 4{

	        A, b := readMatrixAndVector()
	        iterations, systemDescription := metodos.SolveSystem(A, b)

	        fmt.Println(systemDescription)

	        for i, x := range iterations {
		        fmt.Printf("Iteration %d: %v\n", i+1, x)
	        }
        
        }else {
            fmt.Println("Opcion no existe")
    }

}

func imprimirData(data map[int][]float64){
        
    for clave, valores := range data { 
        fmt.Printf("Clave: %d\n", clave) 
        for _, valor := range valores { 
            fmt.Printf("\tValor: %f\n", valor) 
        } 
    }
}


func readMatrixAndVector() ([][]float64, []float64) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the number of rows for the matrix:")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	A := make([][]float64, n)
	fmt.Println("Enter the matrix row by row (space-separated values):")
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := strings.Fields(scanner.Text())
		A[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			A[i][j], _ = strconv.ParseFloat(row[j], 64)
		}
	}

	fmt.Println("Enter the RHS vector (space-separated values):")
	scanner.Scan()
	bInput := strings.Fields(scanner.Text())
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i], _ = strconv.ParseFloat(bInput[i], 64)
	}

	return A, b
}


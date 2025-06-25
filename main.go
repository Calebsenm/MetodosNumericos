package main

import (
	"bufio"
	"fmt"
	"metodosNumericos/metodos"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("---------------- Metodos --------------------")

	var opciones map[string]string = map[string]string{}
	opciones["1"] = "Biseccion"
	opciones["2"] = "Regla Falsa"
	opciones["3"] = "Newthon Raphon"
	opciones["4"] = "Gauss Sediel"

	for a, b := range opciones {
		fmt.Println(a, b)
	}

	var option int
	fmt.Print("Selecciona: ")
	fmt.Scan(&option)
	bufio.NewReader(os.Stdin).ReadString('\n')

	if option == 1 {
		var num1 float64
		fmt.Print("ingresa el num1: ")
		fmt.Scan(&num1)

		fmt.Print("ingresa el num2: ")
		var num2 float64
		fmt.Scan(&num2)

		fmt.Print("ingresa el numero interaciones: ")
		var numIter int
		fmt.Scan(&numIter)

		fmt.Print("ingresa la funcion: ")
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n') // Consumir salto de línea pendiente
		function, _ := reader.ReadString('\n')
		function = strings.TrimSpace(function)

		bs := metodos.Biseccion{NumberOne: num1, NumberTwo: num2, MaxIter: numIter, Function: function}
		data, err := bs.Calcular()
		imprimirData(data)
		if err != nil {
			fmt.Println(err)
		}

	} else if option == 2 {

		var num1 float64
		fmt.Print("ingresa el num1: ")
		fmt.Scan(&num1)

		fmt.Print("ingresa el num2: ")
		var num2 float64
		fmt.Scan(&num2)

		fmt.Print("ingresa el numero interaciones: ")
		var numIter int
		fmt.Scan(&numIter)

		fmt.Print("ingresa la funcion: ")
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n') // Consumir salto de línea pendiente
		function, _ := reader.ReadString('\n')
		function = strings.TrimSpace(function)

		bs := metodos.ReglaFalsa{NumberOne: num1, NumberTwo: num2, MaxIter: numIter, Function: function}
		data, err := bs.Calcular()

		imprimirData(data)

		if err != nil {
			fmt.Println(err)
		}

	} else if option == 3 {

		var num1 float64
		fmt.Print("Ingresa el numero: ")
		fmt.Scan(&num1)

		fmt.Print("ingresa el numero interaciones: ")
		var numIter int
		fmt.Scan(&numIter)

		fmt.Print("ingresa la funcion: ")
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n') // Consumir salto de línea pendiente
		function, _ := reader.ReadString('\n')
		function = strings.TrimSpace(function)

		bs := metodos.NewthonRp{Number: num1, MaxIter: numIter, Function: function}
		result, err := bs.Calular()
		fmt.Println("Aproximaciones de la raíz:")
		for i, v := range result {
			fmt.Printf("Iteración %d: %v\n", i+1, v)
		}
		if err != nil {
			fmt.Println(err)
		}

	} else if option == 4 {

		A, b := readMatrixAndVector()
		iterations, systemDescription := metodos.SolveSystem(A, b)

		fmt.Println(systemDescription)

		for i, x := range iterations {
			fmt.Printf("Iteration %d: %v\n", i+1, x)
		}

	} else {
		fmt.Println("Opcion no existe")
	}

}

func imprimirData(data map[int][]float64) {

	for clave, valores := range data {
		fmt.Printf("Clave: %d\n", clave)
		for _, valor := range valores {
			fmt.Printf("\tValor: %f\n", valor)
		}
	}
}

func readMatrixAndVector() ([][]float64, []float64) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ejemplo de sistema 3x3:")
	fmt.Println("  4 1 2")
	fmt.Println("  3 5 1")
	fmt.Println("  1 1 3")
	fmt.Println("Vector RHS: 4 7 3")
	fmt.Println("Enter the number of rows for the matrix:")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	if n <= 0 {
		fmt.Println("Número de filas inválido.")
		os.Exit(1)
	}

	A := make([][]float64, n)
	fmt.Println("Enter the matrix row by row (space-separated values):")
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := strings.Fields(scanner.Text())
		if len(row) != n {
			fmt.Printf("Fila %d inválida. Debe tener %d valores.\n", i+1, n)
			os.Exit(1)
		}
		A[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			val, err := strconv.ParseFloat(row[j], 64)
			if err != nil {
				fmt.Printf("Valor inválido en la fila %d, columna %d.\n", i+1, j+1)
				os.Exit(1)
			}
			A[i][j] = val
		}
	}

	fmt.Println("Enter the RHS vector (space-separated values):")
	scanner.Scan()
	bInput := strings.Fields(scanner.Text())
	if len(bInput) != n {
		fmt.Println("El vector RHS debe tener el mismo número de valores que filas la matriz.")
		os.Exit(1)
	}
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		val, err := strconv.ParseFloat(bInput[i], 64)
		if err != nil {
			fmt.Printf("Valor inválido en el vector RHS en la posición %d.\n", i+1)
			os.Exit(1)
		}
		b[i] = val
	}

	return A, b
}

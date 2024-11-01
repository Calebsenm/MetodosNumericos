package metodos

import (
    "github.com/Knetic/govaluate"
    "fmt"
    "math"
 
)

// funcion to evalutate the  math expression 
func evaluate(x float64 , function string ) (float64 , error ) {
    expresion, err := govaluate.NewEvaluableExpressionWithFunctions(function , funcionesMatematicas())
	
    if err != nil {
		return 0, fmt.Errorf("error en la expresión: %v", err)
	}
   
	parameters := make(map[string]interface{}, 8)
	parameters["x"] = x ;

	result , err := expresion.Evaluate(parameters)
	if err != nil {
		return 0, fmt.Errorf("error al evaluar la expresión: %v", err)
	}

	return result.(float64), nil
}  

// maths functions for govaluate 
func funcionesMatematicas() map[string]govaluate.ExpressionFunction {
	return map[string]govaluate.ExpressionFunction{
		"sin": func(args ...interface{}) (interface{}, error) {
			return math.Sin(args[0].(float64)), nil
		},
		"cos": func(args ...interface{}) (interface{}, error) {
			return math.Cos(args[0].(float64)), nil
		},
		"tan": func(args ...interface{}) (interface{}, error) {
			return math.Tan(args[0].(float64)), nil
		},
		"asin": func(args ...interface{}) (interface{}, error) {
			return math.Asin(args[0].(float64)), nil
		},
		"acos": func(args ...interface{}) (interface{}, error) {
			return math.Acos(args[0].(float64)), nil
		},
		"atan": func(args ...interface{}) (interface{}, error) {
			return math.Atan(args[0].(float64)), nil
		},
		"sinh": func(args ...interface{}) (interface{}, error) {
			return math.Sinh(args[0].(float64)), nil
		},
		"cosh": func(args ...interface{}) (interface{}, error) {
			return math.Cosh(args[0].(float64)), nil
		},
		"tanh": func(args ...interface{}) (interface{}, error) {
			return math.Tanh(args[0].(float64)), nil
		},
		"sqrt": func(args ...interface{}) (interface{}, error) {
			return math.Sqrt(args[0].(float64)), nil
		},
		"pow": func(args ...interface{}) (interface{}, error) {
			return math.Pow(args[0].(float64), args[1].(float64)), nil
		},
		"ln": func(args ...interface{}) (interface{}, error) {
			return math.Log(args[0].(float64)), nil
		},
		"log10": func(args ...interface{}) (interface{}, error) {
			return math.Log10(args[0].(float64)), nil
		},
		"log": func(args ...interface{}) (interface{}, error) {
			return math.Log(args[0].(float64)) / math.Log(args[1].(float64)), nil
		},
		"exp": func(args ...interface{}) (interface{}, error) {
			return math.Exp(args[0].(float64)), nil
		},
		"round": func(args ...interface{}) (interface{}, error) {
			return math.Round(args[0].(float64)), nil
		},
		"floor": func(args ...interface{}) (interface{}, error) {
			return math.Floor(args[0].(float64)), nil
		},
		"ceil": func(args ...interface{}) (interface{}, error) {
			return math.Ceil(args[0].(float64)), nil
		},
		"abs": func(args ...interface{}) (interface{}, error) {
			return math.Abs(args[0].(float64)), nil
		},
	}
}


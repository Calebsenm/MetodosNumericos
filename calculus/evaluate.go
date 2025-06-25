package calculus

import (
	"fmt"
	"math"
	"strings"

	"github.com/Knetic/govaluate"
)

// funcion to evalutate the  math expression
func EvaluateFunction(x float64, function string) (float64, error) {
	function = insertMultiplication(function)
	function = convertPowers(function)
	expresion, err := govaluate.NewEvaluableExpressionWithFunctions(function, funcionesMatematicas())

	if err != nil {
		return 0, fmt.Errorf("error en la expresión: %v", err)
	}

	parameters := make(map[string]interface{}, 8)
	parameters["x"] = x

	result, err := expresion.Evaluate(parameters)
	if err != nil {
		return 0, fmt.Errorf("error al evaluar la expresión: %v", err)
	}

	return result.(float64), nil
}

// Convierte 'x^n' en 'pow(x,n)' para que govaluate lo entienda
func convertPowers(expr string) string {
	// Reemplaza patrones como 'x^2' o '2^x' por 'pow(x,2)' o 'pow(2,x)'
	var result strings.Builder
	for i := 0; i < len(expr); {
		if i+2 < len(expr) && ((expr[i] == 'x' && expr[i+1] == '^' && (expr[i+2] == 'x' || (expr[i+2] >= '0' && expr[i+2] <= '9'))) || ((expr[i] >= '0' && expr[i] <= '9') && expr[i+1] == '^' && (expr[i+2] == 'x' || (expr[i+2] >= '0' && expr[i+2] <= '9')))) {
			// Caso 'x^n' o 'n^x'
			baseStart := i
			baseEnd := i
			if expr[i] == 'x' {
				baseEnd = i + 1
			} else {
				for baseEnd < len(expr) && expr[baseEnd] >= '0' && expr[baseEnd] <= '9' {
					baseEnd++
				}
			}
			base := expr[baseStart:baseEnd]
			expStart := baseEnd + 1
			expEnd := expStart
			if expr[expStart] == 'x' {
				expEnd = expStart + 1
			} else {
				for expEnd < len(expr) && expr[expEnd] >= '0' && expr[expEnd] <= '9' {
					expEnd++
				}
			}
			exp := expr[expStart:expEnd]
			result.WriteString("pow(" + base + "," + exp + ")")
			i = expEnd
		} else {
			result.WriteByte(expr[i])
			i++
		}
	}
	return result.String()
}

// Inserta '*' entre número y variable, por ejemplo '2x' -> '2*x'
func insertMultiplication(expr string) string {
	var result strings.Builder
	for i := 0; i < len(expr)-1; i++ {
		result.WriteByte(expr[i])
		if expr[i] >= '0' && expr[i] <= '9' && expr[i+1] == 'x' {
			result.WriteByte('*')
		}
	}
	if len(expr) > 0 {
		result.WriteByte(expr[len(expr)-1])
	}
	return result.String()
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

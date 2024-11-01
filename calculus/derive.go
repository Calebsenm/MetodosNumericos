//Este codigo fue generado por ia.
//No fue testeado.
//puede realizar derivadas basicas.

package calculus

import (
	"fmt"
	"strconv"
	"strings"
)

func Derive(expr string) (string, error) {
	if expr == "0" {
		return "(0)", nil
	}

	expr = strings.ReplaceAll(expr, " ", "")
	terms := splitTerms(expr)
	derivatives := make([]string, 0)

	for _, term := range terms {
		derivative, err := deriveTerm(term)
		if err != nil {
			return "", err
		}
		if derivative != "0" {
			derivatives = append(derivatives, derivative)
		}
	}

	if len(derivatives) == 0 {
		return "0", nil
	}

	return simplify(derivatives), nil
}

func splitTerms(expr string) []string {
	var terms []string
	var currentTerm strings.Builder
	var parenCount int

	for i := 0; i < len(expr); i++ {
		c := expr[i]
		if c == '(' {
			parenCount++
		} else if c == ')' {
			parenCount--
		}

		if parenCount == 0 && (c == '+' || (c == '-' && i > 0)) {
			if currentTerm.Len() > 0 {
				terms = append(terms, currentTerm.String())
				currentTerm.Reset()
			}
			if c == '-' {
				currentTerm.WriteRune('-')
			}
		} else {
			currentTerm.WriteByte(c)
		}
	}

	if currentTerm.Len() > 0 {
		terms = append(terms, currentTerm.String())
	}

	return terms
}

func deriveTerm(term string) (string, error) {
	isNegative := false
	if strings.HasPrefix(term, "-") {
		isNegative = true
		term = term[1:]
	}

	// Manejar funciones
	if strings.Contains(term, "(") {
		return deriveFunctionTerm(term, isNegative)
	}

	// Manejar productos
	if strings.Count(term, "*") > 0 {
		return deriveProduct(term, isNegative)
	}

	// Manejar términos polinomiales
	if strings.Contains(term, "x") {
		return derivePolynomial(term, isNegative)
	}

	// Constantes
	if _, err := strconv.ParseFloat(term, 64); err == nil {
		return "0", nil
	}

	return "", fmt.Errorf("término no reconocido: %s", term)
}

func deriveFunctionTerm(term string, isNegative bool) (string, error) {
	if strings.HasPrefix(term, "sin(") {
		inner := term[4 : len(term)-1]
		innerDeriv, err := Derive(inner)
		if err != nil {
			return "", err
		}
		result := "cos(" + inner + ")"
		if innerDeriv != "1" {
			result = fmt.Sprintf("%s*(%s)", result, innerDeriv)
		}
		if isNegative {
			result = "-" + result
		}
		return result, nil
	}

	if strings.HasPrefix(term, "cos(") {
		inner := term[4 : len(term)-1]
		innerDeriv, err := Derive(inner)
		if err != nil {
			return "", err
		}
		result := "-sin(" + inner + ")"
		if innerDeriv != "1" {
			result = fmt.Sprintf("%s*(%s)", result, innerDeriv)
		}
		if isNegative {
			result = "-" + result
		}
		return result, nil
	}

	if strings.HasPrefix(term, "tan(") {
		if isNegative {
			return "-sec^2(x)", nil
		}
		return "sec^2(x)", nil
	}

	if strings.HasPrefix(term, "ln(") {
		if isNegative {
			return "-1/x", nil
		}
		return "1/x", nil
	}

	if strings.HasPrefix(term, "exp(") {
		inner := term[4 : len(term)-1]
		if isNegative {
			return "-exp(" + inner + ")", nil
		}
		return "exp(" + inner + ")", nil
	}

	return "", fmt.Errorf("función no reconocida: %s", term)
}

func deriveProduct(term string, isNegative bool) (string, error) {
	parts := strings.Split(term, "*")

	// Manejar el caso especial de coeficientes numéricos
	if len(parts) == 2 && !strings.Contains(parts[0], "x") {
		coeff, err := strconv.Atoi(parts[0])
		if err == nil {
			deriv, err := deriveTerm(parts[1])
			if err != nil {
				return "", err
			}
			if isNegative {
				coeff = -coeff
			}
			return fmt.Sprintf("%d*%s", coeff, deriv), nil
		}
	}

	// Regla del producto general
	derivatives := make([]string, len(parts))
	for i := range parts {
		product := make([]string, len(parts))
		copy(product, parts)
		deriv, err := deriveTerm(parts[i])
		if err != nil {
			return "", err
		}
		product[i] = deriv
		derivatives[i] = strings.Join(product, "*")
	}

	result := strings.Join(derivatives, " + ")
	if isNegative {
		result = "-(" + result + ")"
	}
	return result, nil
}

func derivePolynomial(term string, isNegative bool) (string, error) {
	if term == "x" {
		if isNegative {
			return "-1", nil
		}
		return "1", nil
	}

	if strings.HasPrefix(term, "x^") {
		exp, err := strconv.Atoi(term[2:])
		if err != nil {
			return "", fmt.Errorf("exponente inválido: %s", term[2:])
		}
		if isNegative {
			return fmt.Sprintf("-%d*x^%d", exp, exp-1), nil
		}
		return fmt.Sprintf("%d*x^%d", exp, exp-1), nil
	}

	// Manejar coeficientes
	if strings.Contains(term, "x^") {
		parts := strings.Split(term, "x^")
		coeff, err := strconv.Atoi(parts[0])
		if err != nil {
			return "", fmt.Errorf("coeficiente inválido: %s", parts[0])
		}
		exp, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("exponente inválido: %s", parts[1])
		}
		newCoeff := coeff * exp
		if isNegative {
			newCoeff = -newCoeff
		}
		if exp-1 == 0 {
			return fmt.Sprintf("%d", newCoeff), nil
		}
		if exp-1 == 1 {
			return fmt.Sprintf("%d*x", newCoeff), nil
		}
		return fmt.Sprintf("%d*x^%d", newCoeff, exp-1), nil
	}

	return "", fmt.Errorf("término polinomial no reconocido: %s", term)
}

func simplify(derivatives []string) string {
	if len(derivatives) == 0 {
		return "0"
	}

	result := derivatives[0]
	for i := 1; i < len(derivatives); i++ {
		if !strings.HasPrefix(derivatives[i], "-") {
			result += " + " + derivatives[i]
		} else {
			result += " " + derivatives[i]
		}
	}

	return result
}

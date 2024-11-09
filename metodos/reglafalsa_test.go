package metodos 

import (
    "testing"
    "fmt"
)

func TestRe(t *testing.T) {
    regla := ReglaFalsa{
        NumberOne: 0,
        NumberTwo: 1,
        MaxIter:   10,
        Function:  "2*x^2 +2*x -2 ",
    }

    data, err := regla.Calcular()
    if err != nil {
        t.Fatalf("Error al calcular: %v", err)
    }

    // Asegúrate de que el mapa data no esté vacío antes de acceder a los valores
    if len(data) == 0 {
        t.Fatal("El mapa data está vacío")
    }

    // Verificar que todas las claves tengan al menos un elemento
    for k, v := range data {
        fmt.Println(k , v )
    }
}

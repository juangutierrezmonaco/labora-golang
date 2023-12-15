package practica

import (
	"fmt"
	"math"
)

func mcm(n1 int, n2 int) int {
	min := int(math.Min(float64(n1), float64(n2)))
	max := n1 * n2
	var mcm int = 1
	for i := min; i <= max; i++ {	
		if math.Mod(float64(i), float64(n1)) == 0 && math.Mod(float64(i), float64(n2)) == 0 {
			mcm = i
			break;
		}
	}
	return mcm
}

func Ej8() {
	// Desarrollar un algoritmo para hallar el mínimo común múltiplo (abreviado como mcm) entre dos números naturales. 
	// El mínimo común múltiplo entre dos números es el menor número que es divisible por ambos.
	// Ej.: mcm (6,9) = 18, mcm (10,15) = 30, mcm (7,14) = 14, mcm (3,7) = 21
	fmt.Print("EJERCICIO 8\n\n")

	var n1, n2 int
	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&n1)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&n2)

	fmt.Printf("El MCM (Mínimo Común Múltiplo) entre %d y %d es: %d", n1, n2, mcm(n1, n2))
}

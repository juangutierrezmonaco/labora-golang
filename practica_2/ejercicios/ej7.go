package practica

import (
	"fmt"
	"math"
)

func mcd(n1 int, n2 int) int {
	nums := []int{n1, n2}
	max := getMaximumNumber(nums)
	var mcd int = 1
	for i := max; i > 0; i-- {
		if math.Mod(float64(n1), float64(i)) == 0 && math.Mod(float64(n2), float64(i)) == 0 {
			mcd = i
			break;
		}
	}
	return mcd
}

func Ej7() {
	// Desarrollar un algoritmo para hallar el máximo común divisor (abreviado como mcd) entre dos números naturales.
	// El máximo común divisor entre dos números es el mayor número que divide a ambos.
	// Ej.: mcd (6,9) = 3, mcd (10,15) = 5, mcd (7,14) = 7, mcd (3,7) = 1
	fmt.Print("EJERCICIO 7\n\n")

	var n1, n2 int
	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&n1)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&n2)

	fmt.Printf("El MCD (Máximo Común Divisor) entre %d y %d es: %d", n1, n2, mcd(n1, n2))
}

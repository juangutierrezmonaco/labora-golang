package practica

import "fmt"

func recursiveNaturalSum(n int) int {
	if n > 0 {
		return n + recursiveNaturalSum((n - 1))
	} else {
		return 0
	}
}


func Ej3() {
	// Realice un algoritmo para sumar los primeros 3 números naturales. Y luego un algoritmo para sumar los primeros 10 números naturales
	fmt.Print("EJERCICIO 3\n\n")

	naturalSumUntil3 := recursiveNaturalSum(3)
	naturalSumUntil10 := recursiveNaturalSum(10)

	fmt.Printf("La suma natural de los primeros 3 términos es: %d\n", naturalSumUntil3)
	fmt.Printf("La suma natural de los primeros 10 términos es: %d\n", naturalSumUntil10)
	
}

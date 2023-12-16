package main

import (
	"fmt"
)

func main() {
	// 	Escribir un algoritmo para sumar los primeros “n” números impares, siendo un “n” un parámetro en un archivo llamado “odds.go”.
	// Luego se desea poner a prueba la correctitud escribiendo en un archivo “odds_test.go”,se pide usar el enfoque de table driven test.
	// Como sugerencia se puede pensar en que cada test puede describirse como:
	// 		type Test struct {
	// 				n int,            // cantidad de primeros impares a sumar
	// 				expectedSum int,  // resultado esperado
	// 		}
	fmt.Print("EJERCICIO 1\n\n")
}

func sumOddOrEvenNumbers(n int, isOdd bool) int {
	acc := 0

	for i := 0; i <= n; i++ {
		acc += 2 * i
	}

	if isOdd {
		// Note: This is because the n term for odd sums is: 2n - 1 (instead of 2n for even numbers)
		// So, if we sum all the even numbers, for every one, there's a -1 that we aren't applying
		// To compensate -1 n times is equal to -n, so we substract from the accumulator.
		acc -= n
	}
	return acc
}

func SumNOddNumbers(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("No se puede sumar una cantidad negativa de enteros")
	}
	return sumOddOrEvenNumbers(n, true), nil
}

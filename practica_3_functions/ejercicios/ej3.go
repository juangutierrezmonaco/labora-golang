package practica

import (
	"fmt"
)

func IsPrimeNumber(num int) bool {
	divs := findNumberDivisors(num)
	return len(divs) == 1
}

func Ej3() {
	// Escriba un algoritmo para determinar si un número es primo. Recordar que número primo es aquel que es divisible por solo por 1 y por si mismo.
	// ¿Se puede utilizar alguna función desarrollada en los ejercicios anteriores?
	fmt.Print("EJERCICIO 3\n\n")

	var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)

	isPrimeNumber := IsPrimeNumber(num)
	if isPrimeNumber {
		fmt.Printf("El número %d es primo", num)
	} else {
		fmt.Printf("El número %d NO es primo", num)
	}
}

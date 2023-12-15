package practica

import (
	"fmt"
)

func div(n, d float64) float64 {
	if d == 0 {
		panic("No se puede dividir por 0.")
	}
	return n / d
}

func recoverAndPrintError() {
	r := recover()
	if r != nil {
		fmt.Println("Hubo un error:", r)
	}
}

func Ej1() {
	// Escribe un programa en Go que solicite al usuario dos números (numerador y denominador), intente realizar la división y maneje el caso en el que el
	// denominador sea cero. Si el denominador es cero, imprime un mensaje de error indicando que la división no es posible.
	// En caso contrario, imprime el resultado de la división con dos decimales.
	fmt.Print("EJERCICIO 1\n\n")

	defer recoverAndPrintError()

	var num, den float64
	fmt.Print("Ingrese un numerador: ")
	fmt.Scan(&num)
	fmt.Print("Ingrese un denominador: ")
	fmt.Scan(&den)

	res := div(num, den)
	fmt.Printf("El resultado de dividir %.2f por %.2f es: %.2f\n", num, den, res)
}

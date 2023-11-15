package practica

import "fmt"

func Ej2() {
	// Escribir un algoritmo que determine si un número es par
	fmt.Print("EJERCICIO 2\n\n")

	var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)

	isEven := num%2 == 0
	var isEvenText string
	if isEven {
		isEvenText = "par"
	} else {
		isEvenText = "impar"
	}
	fmt.Println("Su número es", num, "y es", isEvenText)
}

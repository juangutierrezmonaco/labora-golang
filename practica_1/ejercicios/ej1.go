package practica

import "fmt"

func Ej1() {
	// Realizar un algoritmo para leer un número e informar si es mayor, igual o menos a cero.
	fmt.Print("EJERCICIO 1\n\n")

	var num int
	var sign string
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)

	switch true {
	case num > 0:
		sign = "mayor"
	case num < 0:
		sign = "menor"
	case num == 0:
		sign = "igual"
	}

	fmt.Println("Su número es", num, "y es", sign, "a 0.")
}

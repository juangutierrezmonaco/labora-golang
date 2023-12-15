package practica

import (
	"fmt"
)

// Nota: En inglés se usa palindromic, se refieren a lo mismo si hablan de números o palabras.
func isCapicua(num int) bool {
	return isPalindromic(fmt.Sprint(num))
}

func Ej9() {
	// Realice un algoritmo que dado un número te diga si es capicua.
	fmt.Print("EJERCICIO 9\n\n")

	// Es exactamente lo mismo que el ejercicio 8, palíndromo y capicúa son dos maneras de decir lo mismo, con la diferencia de que si son números se le dice capicúa.
	// Pero la lógica es idéntica.

	fmt.Print("Ingrese un número para saber si es capicúa: ")
	var num int
	fmt.Scan(&num)

	if isCapicua(num) {
		fmt.Printf("El número %d es capicúa.\n", num)
	} else {
		fmt.Printf("El número %d NO es capicúa.\n", num)
	}
}

package practica

import "fmt"


func naturalSumGauss(n int) int {
	return (n * (n + 1)) / 2
}

func Ej4() {
	// Realice un algoritmo para sumar los primeros "n" números naturales, siendo "n" un valor que ingresa el usuario por teclado durante la ejecución del algoritmo
	var num int
	fmt.Print("Ingrese el número hasta el qué desea sumar: ")
	fmt.Scan(&num)

	fmt.Printf("La suma natural de los primeros %d términos es: %d\n", num, naturalSumGauss(num))
}

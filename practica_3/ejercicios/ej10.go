package practica

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func sum(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + sum(n-1)
	}
}

func factorialTimesSum(n int) int {
	factorial, sum := 1, 0
	for i := n; i > 0; i-- {
		factorial *= i
		sum += i
	}
	return factorial * sum
}

func Ej10() {
	// Realice un algoritmo para multiplicar el factorial de un número por su sumatoria.
	fmt.Print("EJERCICIO 10\n\n")

	var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)

	fmt.Println(factorial(num) * sum(num))

	// ó

	fmt.Println(factorialTimesSum(num))
}

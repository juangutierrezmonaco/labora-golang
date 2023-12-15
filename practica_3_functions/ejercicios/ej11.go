package practica

import (
	"fmt"
)

func pow(base int, exponent int) int {
	pow := 1
	for i := 0; i < exponent; i++ {
		pow *= base
	}
	return pow
}

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}

func mcm(n1 int, n2 int) int {
	min := min(n1, n2)
	max := n1 * n2
	var mcm int = 1
	for i := min; i <= max; i++ {
		if i%n1 == 0 && i%n2 == 0 {
			mcm = i
			break
		}
	}
	return mcm
}

func vituFunc(x int, y int) (float64, []int) {
	t1 := pow(x, y)
	t2 := sum(x) * y
	t3 := mcm(x, y)

	terms := []int{t1, t2, t3}
	return float64(t1+t2) / float64(t3), terms
}

func Ej11() {
	// Realice un algoritmo que dado dos números calcule el resultado de la potencia del primero elevado al segundo
	// más la sumatoria del primero multiplicado el segundo, todo lo anterior dividido el mínimo común múltiplo entre ambos números.

	//	 x^y + sum(1, x)*y
	// 	 -----------------
	//	     mcm(x, y)

	// No soy una mala persona. Es un lindo ejercicio para aprovechar y modularizar.
	// Vale suponer que x, y son números distintos de cero.
	// Recuerde que la potencia de un número “x” a la “y” consiste en multiplicar “y” veces al número “x” por si mismo.
	fmt.Print("EJERCICIO 11\n\n")

	var n1, n2 int
	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&n1)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&n2)

	res, terms := vituFunc(n1, n2)

	fmt.Printf("El resultado de la función en (%d, %d) es: \n\n", n1, n2)
	fmt.Printf("x^y + sum(1, x)*y   =    %d * %d \n", terms[0], terms[1])
	fmt.Printf("-----------------   =   ---------\n")
	fmt.Printf("    mcm(x, y)       =      %d    \n", terms[2])
	fmt.Println()

	fmt.Printf("Y el resultado es: %.2f\n", res)
}

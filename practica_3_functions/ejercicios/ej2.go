package practica

import (
	"fmt"

	"github.com/labora/labora-golang/util"
)

func areFriendNumbers(n1 int, n2 int) (bool, []int, []int) {
	_, divs1, s1 := isPerfectNumber(n1)
	_, divs2, s2 := isPerfectNumber(n2)
	return n1 == s2 && n2 == s1, divs1, divs2
}

func Ej2() {
	// Proponga una solución que use funciones bien diseñadas (solución modular) al problema de determinar si dos números naturales son amigos. O sea la ideas es pensar en cómo puede reutilizar la función del ejercicio anterior para si un número es perfecto.
	// Un número es amigo de otro cuando la suma de sus divisores propios es igual al otro número.
	// Sean X1, X2, XN todos divisores propios de X
	// Sean Y1, Y2, YN todos divisores propios de Y
	// Si X e Y son amigos entonces X1 + x2 +…+ XN es igual a Y e Y1 + Y2 +…+ XN es igual a X
	// Ejemplo:
	// El par (220, 284), ya que:
	// Los divisores propios de 220 son 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 y 110, que suman 284.
	// Los divisores propios de 284 son 1, 2, 4, 71 y 142, que suman 220.
	fmt.Print("EJERCICIO 2\n\n")

	var n1, n2 int
	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&n1)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&n2)

	areFriendNumbers, d1, d2 := areFriendNumbers(n1, n2)
	joinedDivisors1 := util.JoinIntArr(d1, " + ")
	joinedDivisors2 := util.JoinIntArr(d2, " + ")

	if areFriendNumbers {
		fmt.Printf("Los números %d y %d son amigos porque:\n", n1, n2)
		fmt.Printf("\t-> %s (divisores de %d) = %d\n", joinedDivisors2, n2, n1)
		fmt.Printf("\t-> %s (divisores de %d) = %d\n", joinedDivisors1, n1, n2)
	} else {
		fmt.Printf("Los números %d y %d NO son amigos porque:\n", n1, n2)
		fmt.Printf("\t-> %s (divisores de %d) ≠ %d\n", joinedDivisors2, n2, n1)
		fmt.Printf("\t-> %s (divisores de %d) ≠ %d\n", joinedDivisors1, n1, n2)
	}
}

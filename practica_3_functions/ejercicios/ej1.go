package practica

import (
	"fmt"

	"github.com/labora/labora-golang/util"
)

func findNumberDivisors(num int) []int {
	var divisors []int
	for i := 1; i < num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func isPerfectNumber(num int) (bool, []int, int) {
	sum := 0
	divisors := findNumberDivisors(num)
	for _, div := range divisors {
		sum += div
	}
	return sum == num, divisors, sum
}

func Ej1() {
	// Desarrolle una función para determinar si un número natural (> 0) es perfecto. Un número es perfecto si la suma de los divisores propios da el el mismo numero.
	// Los divisores propios de un número son todos sus divisores sin contar el mismo número.
	// Sean X1, X2, XN todos divisores propios de X
	// Si X es propio entonces X1 + x2 +…+ XN es igual a X
	// Ejemplo:
	// 6 es un número perfecto
	// Divisores Propios: 1, 2 y 3.
	// 1 + 2 + 3 = 6
	fmt.Print("EJERCICIO 1\n\n")

	var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)

	isPerfectNumber, divs, _ := isPerfectNumber(num)
	joinedDivisors := util.JoinIntArr(divs, "+")
	if isPerfectNumber {
		fmt.Printf("El número %d es perfecto, pues %s = %d\n", num, joinedDivisors, num)
	} else {
		fmt.Printf("El número %d no es perfecto, pues %s ≠ %d\n", num, joinedDivisors, num)
	}
}

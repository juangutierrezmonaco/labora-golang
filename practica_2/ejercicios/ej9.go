package practica

import (
	"fmt"
	"math"
	"strconv"
)

func numberToDigitsArray(num int) []int {
	var arr []int
	numStr := strconv.Itoa(num)
	for i := 0; i < len(numStr); i++ {
		n, _ := strconv.Atoi(string(numStr[i]))
		arr = append(arr, n)
	}
	return arr
}

func isPalindromicNumber(num int) bool {
	numArr := numberToDigitsArray(num)
	middleIndex, _ := math.Modf(float64(len(numArr) / 2))
		
	// Si es impar, no va a considerar el elemento del medio. Si es par, sí.
	// Por ej: 	12345 => middleIndex = int(2,5) = 2 ==> No lo considera
	// 				  123456 => middleIndex = int(3) = 3 ==> Lo considera

	for i := 0; i < int(middleIndex); i++ {
		// El índice i va a recorrer el arreglo de manera progresiva y el j, regresiva
		j := len(numArr) - 1 - i
		if numArr[i] != numArr[j] {
			return false
		}
	}

	return true
}

func Ej9() {
	// Dado un número de 5 cifras, determinar si es capicúa.
	// Si fuera un número de 6 cifras ¿Sirve la resolución planteada? ¿Cómo habría que modificarla?
	fmt.Print("EJERCICIO 9\n\n")

	var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)

	if isPalindromicNumberB(num) {
		fmt.Printf("El número %d es capicúa.", num)
	} else {
		fmt.Printf("El número %d NO es capicúa.", num)
	}
}

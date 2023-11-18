package practica

import (
	"fmt"
	"math"

	"github.com/labora/labora-golang/stack"
)

func removeMiddleDigitIfNecessary(num int) []int {
	numArr := numberToDigitsArray(num)
	digitsSize := len(numArr)
	if math.Mod(float64(digitsSize), 2) != 0 {
		middleIndex, _ := math.Modf(float64(digitsSize / 2))
		numArr = append(numArr[:int(middleIndex)], numArr[int(middleIndex)+1:]...)
	}
	return numArr
}

func isPalindromicNumberB(num int) bool {
	// El dígito del medio no afecta si el número es capicúa cuando éste tiene una cantidad de dígitos impar
	numArr := removeMiddleDigitIfNecessary(num)

	s := stack.NewStack()
	s.Push(numArr[0])
	for i := 1; i < len(numArr); i++ {
		digit := numArr[i]

		if s.Peek() == digit {
			// Si está en la pila, removelo
			s.Pop()
		} else {
			// Sino, agregalo
			s.Push(digit)
		}
	}
	// Si era capicúa, la pila tendría que estar vacía
	return s.IsEmpty()
}

func Ej9B() {
	// Dado un número de 5 cifras, determinar si es capicúa.
	// Si fuera un número de 6 cifras ¿Sirve la resolución planteada? ¿Cómo habría que modificarla?
	fmt.Print("EJERCICIO 9\n\n")

	var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)

	fmt.Println()

	if isPalindromicNumber(num) {
		fmt.Printf("El número %d es capicúa.", num)
	} else {
		fmt.Printf("El número %d NO es capicúa.", num)
	}
}

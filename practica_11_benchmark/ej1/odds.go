package main

import (
	"fmt"
)

func SumNOddNumbersLoop1(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("No se puede sumar una cantidad negativa de enteros")
	}
	acc := 0

	for i := 1; i <= n; i++ {
		acc += 2*i - 1
	}
	return acc, nil
}

func SumNOddNumbersLoop2(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("No se puede sumar una cantidad negativa de enteros")
	}
	acc := 0
	oddQuantity := 0

	for i := 0; oddQuantity < n; i++ {
		if i%2 != 0 {
			acc += i
			oddQuantity++
		}
	}
	return acc, nil
}

func SumNOddNumbersGauss(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("No se puede sumar una cantidad negativa de enteros")
	}
	sum := n*n
	return sum, nil
}

func main() {
	// Se pide implementar al menos dos formas de sumar los primeros “n” números impares (una de ellas deberías tener por el ejercicio 1 de esta guía). Y luego comparar cual es más eficiente.
	// Podés guiarte con este código(https://github.com/vituchon/labora-golang-course/blob/master/count-pairs-benchmark/main_test.go) (Ver `func Benchmark_sumPair`)
	fmt.Print("EJERCICIO 1\n\n")
}

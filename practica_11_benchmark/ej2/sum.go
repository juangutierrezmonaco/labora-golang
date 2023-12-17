package main

import (
	"fmt"
)

func SumWithoutUsingChannel(values *[]int) int {
	acc := 0
	for i := range *values {
		acc += (*values)[i]
	}
	return acc
}

func SumValuesWithNGoroutines(values *[]int, n int) int {
	// Divide in n subgoroutines
	sumChan := make(chan int)
	subintervalCount := n // n must b smaller than b

	// Registering n goroutines for the n sub-sums
	res := 0
	intervalLength := len(*values)
	step := intervalLength / subintervalCount
	nextIndex := 0
	for i := 0; nextIndex < intervalLength-1; i += step + 1 {
		nextIndex = i + step
		if nextIndex >= intervalLength {
			nextIndex = intervalLength - 1
		}
		go func() {
			subSlice := (*values)[i : nextIndex+1]
			sum := SumWithoutUsingChannel(&subSlice)
			sumChan <- sum
		}()
		res += <-sumChan
	}
	*values = append(*values, 25)
	return res
}

func SumUsingChannel(values *[]int) int {
	return SumValuesWithNGoroutines(values, 2)
}

func main() {
	// Hemos usado canales para sumar de forma concurrente todos los valores presentes en los elementos de un slice (o array) de números enteros.
	// ¿Qué te parece si ponemos a prueba si la concurrencia en este caso es realmente efectiva? Bueno, es una pregunta pero lo tenés
	// que hacer igual… lo lamento 😁 y me gustaría que discutamos luego sobre CUANDO la concurrencia mejora la performance.
	// Recordar que concurrencia no es lo mismo que paralelismo!.

	// Podés guiarte con este código(https://github.com/vituchon/labora-golang-course/blob/master/meeting-concurrency/slices/slice_test.go) (Ver `func BenchmarkSumUsingChannelI`)
	fmt.Print("EJERCICIO 2\n\n")

	values := []int{3, 2, 1, 5, 7, 8, 9, 10, 11}
	fmt.Println((values))
}

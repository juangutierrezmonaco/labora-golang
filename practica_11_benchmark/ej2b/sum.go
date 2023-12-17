package main

import (
	"fmt"

	"github.com/labora/labora-golang/util"
)

func sumValues(values *[]int) int {
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
			sum := sumValues(&subSlice)
			sumChan <- sum
		}()
		res += <-sumChan
	}

	return res
}

func main() {
	// Ahora voy a comparar la eficiencia diviendo el intervalo de suma en N gorutinas
	fmt.Print("EJERCICIO 2b\n\n")

	values := []int{3, 2, 1, 5, 7, 8, 9, 10, 11}
	numberOfGoroutines := util.GenerateRandomNumber(1, len(values))
	res := SumValuesWithNGoroutines(&values, len(values))
	fmt.Printf("El resultado de sumar los valores de %v, usando %d gorutinas, es: %d\n", values, numberOfGoroutines, res)
}

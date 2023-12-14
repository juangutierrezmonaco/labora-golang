package practica

import (
	"fmt"
	"time"
)

func timer(ch *chan time.Duration) {
	start := time.Now()
	<-*ch
	*ch <- time.Since(start)
}

func naturalSum(from int, to int) int {
	acc := 0
	for i := from; i <= to; i++ {
		acc += i
	}
	return acc
}

func naturalSumWithConcurrence(from int, to int, ch *chan int) {
	sum := naturalSum(from, to)
	*ch <- sum
}

func Ej1() {
	// Hacer una suma de los primeros 100 numeros, de forma secuencial y luego de forma concurrente y apreciar la diferencia en el tiempo de ejecución.
	fmt.Print("EJERCICIO 1\n\n")

	timerChan := make(chan time.Duration)
	from, to := 1, 100

	// Sequential
	go timer(&timerChan)

	res := naturalSum(from, to)
	fmt.Println("El resultado (secuencial) es: ", res)

	timerChan <- 0
	fmt.Printf("Tardó: %v\n", <-timerChan)

	// Concurre	nce
	go timer(&timerChan)

	// Divide in n subgoroutines
	sumChan := make(chan int)
	n := 30 // n must b smaller than b

	intervalLength := to - from + 1
	step := intervalLength / n
	an, bn := 0, 0

	// Registering n goroutines for the n sub-sums
	res = 0
	for i := 0; bn < to; i++ {
		an = (step * i) + 1
		bn = step * (i + 1)
		if bn > to {
			bn = to
		}
		go naturalSumWithConcurrence(an, bn, &sumChan)
		res += <-sumChan
	}
	fmt.Println("El resultado (concurrente) es: ", res)

	timerChan <- 0 // Send to the timer that it finished
	fmt.Printf("Tardó: %v\n", <-timerChan)
}

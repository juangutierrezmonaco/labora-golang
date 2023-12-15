package practica

import (
	"fmt"
	"sync"
	"time"
)

func sumOddOrEvenNumbers(n int, isOdd bool) int {
	acc := 0

	for i := 0; i <= n; i++ {
		acc += 2 * i
	}

	if isOdd {
		// Note: This is because the n term for odd sums is: 2n - 1 (instead of 2n for even numbers)
		// So, if we sum all the even numbers, for every one, there's a -1 that we aren't applying
		// To compensate -1 n times is equal to -n, so we substract from the accumulator.
		acc -= n
	}
	return acc
}

func sumNEvenNumbers(n int) int {
	return sumOddOrEvenNumbers(n, false)
}

func sumNEvenNumbersGauss(n int) int {
	return n * (n + 1)
}

func sumNOddNumbers(n int) int {
	return sumOddOrEvenNumbers(n, true)
}

func sumNOddNumbersGauss(n int) int {
	return n * n
}

func concurrentSum(n int, numberText string, wg *sync.WaitGroup, sumFunc func(int) int) {
	timerChan := make(chan time.Duration)
	defer wg.Done()

	go timer(&timerChan)
	res := sumFunc(n)
	timerChan <- 0
	fmt.Printf("La suma de los primeros %d números %s dio: %d. Y tardó: %v\n", n, numberText, res, <-timerChan)
}

func Ej3() {
	// Escriba un programa en donde haya dos gorutinas donde una suma los primeros 100 números pares y la otra los 100 primeros números impares.
	// Se desea saber cual gorutina “gana”, es decir, quien termina termina primero.
	// Para sincronizar se puede usar grupos de espera o canales. Preferentemente podés realizar el ejercicio con ambos mecanismos para que te vayas familiarizando.
	fmt.Print("EJERCICIO 3\n\n")

	wg := sync.WaitGroup{}
	wg.Add(2)
	n := 10000

	// Even sum
	go concurrentSum(n, "pares", &wg, sumNEvenNumbers)

	// Odd sum
	go concurrentSum(n, "impares", &wg, sumNOddNumbers)

	wg.Wait()
}

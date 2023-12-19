package practica

import (
	"fmt"
)

func Ej3b() {
	// Desarrolla un programa en Golang con un canal (de enteros) compartido entre dos hilos productores y un hilo consumidor.
	// La regla es que los hilos productores deben colocar valores en el canal de manera intercalada, asegurándose de que un hilo productor
	// no ponga dos valores consecutivos sin que el otro hilo productor haya puesto uno.
	// Generar la secuencia de n números naturales, mandando los números pares por un lado, y los impares por el otro.
	fmt.Print("EJERCICIO 3\n\n")

	valuesCh := make(chan int)
	signalCh := make(chan bool)
	exitCh := make(chan bool)
	n := 100
	exit := false

	count := 0

	// producer of even numbers
	go func() {
		for !exit {
			nextEvenNumber := 2 * count

			fmt.Println("[ PAR ] - Mandando número: ", nextEvenNumber)
			valuesCh <- nextEvenNumber

			fmt.Println("[ PAR ] - Mandando señal para que genere el otro hilo...")
			signalCh <- true

			<-signalCh
			fmt.Println("[ PAR ] - Esperando para generar próximo número...")
		}
	}()

	// producer of odd numbers
	go func() {
		for !exit {
			nextOddNumber := 2*count + 1

			<-signalCh
			fmt.Println("[ IMPAR ] - Esperando para generar próximo número...")

			fmt.Println("[ IMPAR ] - Mandando número: ", nextOddNumber)
			valuesCh <- nextOddNumber

			count++
			fmt.Println("[ IMPAR ] - Mandando señal para que genere el otro hilo...")
			signalCh <- true
		}
	}()

	// consumer
	go func() {
		var nums []int
		for i := 0; i < n; i++ {
			nums = append(nums, <-valuesCh)
		}
		fmt.Printf("\n\nEl resultado es: %v\n\n", nums)
		exitCh <- true
	}()

	<-exitCh
	exit = true
}

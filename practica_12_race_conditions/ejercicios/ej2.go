package practica

import (
	"fmt"
	"sync"
)

func Ej2() {
	// Desarrolla un programa en Golang con un canal (de enteros) compartido entre dos hilos productores y un hilo consumidor.
	// La regla es que los hilos productores deben colocar valores en el canal de manera intercalada, asegurándose de que un hilo productor
	// no ponga dos valores consecutivos sin que el otro hilo productor haya puesto uno.
	fmt.Print("EJERCICIO 2\n\n")

	valuesCh := make(chan int)
	signalCh := make(chan bool)
	exitCh := make(chan bool)
	n := 10
	exit := false

	// producer of ceros
	go func() {
		for !exit {
			valuesCh <- 0
			signalCh <- true
			<-signalCh
		}
	}()

	// producer of ones
	go func() {
		for !exit {
			<-signalCh
			valuesCh <- 1
			signalCh <- true
		}
	}()

	// consumer
	go func() {
		for i := 0; i < n; i++ {
			fmt.Printf("Número %d: %d\n", i, <-valuesCh)
		}
		exitCh <- true
	}()

	<-exitCh
	exit = true
}

// POR QUÉ NO ANDA CON WAITGROUPS
// HAY QUE CERRAR LOS CANALES?
func Ej2WG() {
	// Desarrolla un programa en Golang con un canal (de enteros) compartido entre dos hilos productores y un hilo consumidor.
	// La regla es que los hilos productores deben colocar valores en el canal de manera intercalada, asegurándose de que un hilo productor
	// no ponga dos valores consecutivos sin que el otro hilo productor haya puesto uno.
	fmt.Print("EJERCICIO 2\n\n")

	valuesCh := make(chan int)
	signalCh := make(chan bool)
	count := 2
	exit := false

	wg := sync.WaitGroup{}
	wg.Add(3)

	// producer of ceros
	go func() {
		defer wg.Done()
		for !exit {
			valuesCh <- 0
			signalCh <- true
			<-signalCh
		}
	}()

	// producer of ones
	go func() {
		defer wg.Done()
		for !exit {
			<-signalCh
			valuesCh <- 1
			signalCh <- true
		}
	}()

	// consumer
	go func() {
		defer wg.Done()
		for i := 0; i < 2*count; i++ {
			fmt.Printf("Número %d: %d\n", i, <-valuesCh)
		}
	}()
	wg.Wait()
	exit = true
}

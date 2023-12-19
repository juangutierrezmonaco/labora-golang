package practica

import (
	"fmt"
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
		fmt.Println("SALIENDO DE PRODUCTOR 1")
	}()

	// producer of ones
	go func() {
		for !exit {
			<-signalCh
			valuesCh <- 1
			signalCh <- true
		}
		fmt.Println("SALIENDO DE PRODUCTOR 2")
	}()

	// consumer
	go func() {
		for i := 0; i < n; i++ {
			fmt.Printf("Número %d: %d\n", i, <-valuesCh)
		}
		exitCh <- true
		fmt.Println("SALIENDO DE CONSUMIDOR")
	}()

	<-exitCh
	exit = true
	// El programa no termina bien, pues uno o los dos productores que dan esperando y sólo se resuelve bien porque se termina el programa.
}

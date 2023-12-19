package practica

import (
	"fmt"
	"sync"

	"github.com/labora/labora-golang/util"
)

func reverseString(str string) string {
	var reverseStr string
	strCh := make(chan rune)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer close(strCh) // Triggers the print goroutine
		for _, c := range str {
			defer func(c rune) {
				strCh <- c
			}(c)
		}
	}()

	go func() {
		defer wg.Done() // Triggers the main to finish
		for c := range strCh {
			reverseStr += string(c)
		}
	}()

	wg.Wait()
	return reverseStr
}

func Ej2() {
	// 	Hace algún tiempo una alumna me acercó una forma interesante de invertir los caracteres de un texto usando la función defer, fue algo como
	// [esto](https://go.dev/play/p/5qDd82o-qW5). Si se quiere tenerlo dentro de una función aparte, se debe recurrir a los *named return parameters*
	// (Tal como hago [acá](https://go.dev/play/p/9SUybBldinA)). Basta de charla, vamos con el ejercicio!.

	// Lo que quiero que hagan es un programita que use canales para enviar desde un hilo los caracteres en el orden inverso,
	// usando la función defer, y en el otro hilo ir recibiendolos para terminar armando el string invertido.

	// [Solución posible](https://go.dev/play/p/_xZGkXMYW5W)
	fmt.Print("EJERCICIO 2\n\n")

	fmt.Print("Ingrese un texto para revertir: ")
	str := util.ScanSentence()
	fmt.Printf("El texto \"%s\" invertido es: \"%s\"\n", str, reverseString(str))

}

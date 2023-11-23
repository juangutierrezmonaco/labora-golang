package practica

import (
	"fmt"

	"github.com/labora/labora-golang/util"
)

func reverseSentence(word string) string {
	reverseSentence := ""
	for i := 0; i < len(word); i++ {
		reverseSentence = reverseSentence + string(word[len(word)-1-i])
	}
	return reverseSentence
}

func Ej7() {
	// Realice un algoritmo que dado un string lo invierta.
	fmt.Print("EJERCICIO 7\n\n")

	fmt.Print("Ingrese una palabra o frase: ")
	sentence := util.ScanSentence()

	fmt.Printf("La frase %s invertida es: %s", sentence, reverseSentence(sentence))
}

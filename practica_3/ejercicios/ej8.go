package practica

import (
	"fmt"

	"github.com/labora/labora-golang/util"
)

func isPalindromic(sentece string) bool {
	return sentece == reverseSentence(sentece)
}

func Ej8() {
	// Realice un algoritmo que dado un string te diga si es palindromo
	fmt.Print("EJERCICIO 8\n\n")

	fmt.Print("Ingrese una palabra o una frase para saber si es palíndroma: ")
	sentence := util.ScanSentence()

	if isPalindromic(sentence) {
		fmt.Printf("La palabra %s es palíndroma.\n", sentence)
	} else {
		fmt.Printf("La palabra %s NO es palíndroma.\n", sentence)
	}

}

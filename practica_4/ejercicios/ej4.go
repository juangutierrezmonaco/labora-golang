package practica

import (
	"fmt"
	"strings"

	"github.com/labora/labora-golang/util"
)

func getEncodedVowel(r rune) rune {
	switch r {
	case 'a':
		return '4'
	case 'e':
		return '3'
	case 'i':
		return '1'
	case 'o':
		return '0'
	case 'u':
		return '6'
	default:
		return r
	}
}

func encodeString(text string) string {
	return strings.Map(getEncodedVowel, text)
}

func Ej4() {
	// Desarrolle una función que reciba como parámetro una cadena y que reemplace cada vocal por un carácter que la represente simbólicamente.
	// Se puede usar la siguiente tabla de transformación:
	// ‘a’ → ‘4’
	// ‘e’ → ‘3’
	// ‘i’ → ‘1’
	// ‘o’ → ‘0’
	// ‘u’ → ‘6’ (no hay mejor candidato, pero si quieren usar otro sean bienvenidos)

	// Asumiendo que un usuario ingresa “pepa” y quiere transformar su cadena usando la codificación correspondiente, quedaría “p3p4”.
	fmt.Print("EJERCICIO 4\n\n")
	
	fmt.Print("Ingrese una cadena de texto: ")
	text := util.ScanSentence()

	encodedText := encodeString(text)
	fmt.Printf("La cadena \"%s\" codificada es igual a: \"%s\"\n", text, encodedText)
}

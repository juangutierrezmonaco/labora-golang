package practica

import (
	"fmt"
	"strings"
)

func cleanString(s string) string {
	cleanedStr := strings.Replace(s, ".", "", -1)
	cleanedStr = strings.Replace(cleanedStr, ",", "", -1)
	return cleanedStr
}

func WordCount(s string) map[string]int {
	countByChar := make(map[string]int)
	wordArr := strings.Fields(cleanString(s))
	for _, word := range wordArr {
		countByChar[word]++
	}
	return countByChar
}

func Ej1() {
	// - Implementar ContadorDePalabras.
	// - Debería devolver un mapa de los recuentos de cada "palabra" en la cadena s.
	// Por ejemplo: `I ate a donut. Then I ate another donut.`
	// Debería devolver: `map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}`

	// Podes encontrar strings.Fields útil para el primer ejercicio.
	// La función strings.Fields() en Golang se usa para dividir la cadena dada alrededor de cada instancia de uno o más caracteres de espacio en blanco consecutivos,
	// como lo define unicode.IsSpace, devolviendo una porción de subcadenas de str o una porción vacía si str contiene solo blanco espacio.
	fmt.Print("EJERCICIO 1\n\n")

	s := "I ate a donut. Then I ate another donut."
	countByChar := WordCount(s)
	fmt.Printf("%v\n", countByChar)
}

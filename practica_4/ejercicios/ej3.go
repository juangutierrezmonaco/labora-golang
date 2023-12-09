package practica

import (
	"fmt"
)

func contarLetra(letra rune, texto string) int {
	cant := 0
	longTexto := len(texto)

	for i := 0; i < longTexto; i++ {
		c := rune(texto[i])
		if c == letra {
			cant++
		}
	}

	return cant
}

func contarVocales1(text string) {
	var cantA, cantE, cantI, cantO, cantU int

	cantA = contarLetra('a', text)
	cantE = contarLetra('e', text)
	cantI = contarLetra('i', text)
	cantO = contarLetra('o', text)
	cantU = contarLetra('u', text)

	fmt.Printf("Repeticiones de a,e,i,o,u son: %d, %d, %d, %d, %d\n", cantA, cantE, cantI, cantO, cantU)
}

func contarVocales2(text string) {
	var cantA, cantE, cantI, cantO, cantU int
	var c rune

	lenText := len(text)
	i := 0

	for i < lenText {
		c = rune(text[i])

		switch c {
		case 'a':
			cantA++
		case 'e':
			cantE++
		case 'i':
			cantI++
		case 'o':
			cantO++
		case 'u':
			cantU++
		}

		i++
	}

	fmt.Printf("Repeticiones de a,e,i,o,u son: %d, %d, %d, %d, %d\n", cantA, cantE, cantI, cantO, cantU)
}

func Ej3() {
	// Para contar la cantidad de vocales que se repiten en texto se provén estos dos algoritmos. Nos gustaría que observes ambas formas de resolver el problema, 
	// entienda los algoritmos y discuta con sus compañeros de la ventaja y desventajas de los enfoques.
	fmt.Print("EJERCICIO 3\n\n")

	//   Forma 1
	contarVocales1("algo lindo")
	//   Forma 2	
	contarVocales2("algo lindo")

	// En contarVocales1 se está recorriendo todo el texto por cada letra, por lo cual, es más optima la segunda manera. Es decir, si tuviéramos
	// un texto con muchísimas letras, se estaría recorriendo el mismo 5 veces.
}

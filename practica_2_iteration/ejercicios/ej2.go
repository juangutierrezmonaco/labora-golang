package practica

import "fmt"

func Ej2() {
	// Detecte el “pequeño inconveniente” con el siguiente algoritmo (no vale correrlo, sino que tienen que pensarlo!
	fmt.Print("EJERCICIO 2\n\n")

	var n int = 0
	for n != 1 || n != 0 {
		n++
	}

	// NUNCA ENTRA, LA CONDICIÓN ES SIEMPRE FALSA.
	// Un número SIEMPRE va a ser distinto de 0 o distinto de 1. Los casos límite serían:
	// * Cuando n = 0: n != 1 es verdadero
	// * Cuando n = 1: n != 0 es verdadero
	// Y en el resto de los casos, siempre van a ser las dos condiciones verdaderas.

	// Otra forma de verlo, es mirar la negación de la condición:
	// n != 1 || n != 0			=			n == 1 && n == 0
	// Si prestamos atención a (n == 1 && n == 0), un número nunca puede ser 0 y 1 al mismo tiempo (a menos que estemos en el mundo de la Mecánica Cuántica, jé)
	// Por lo cual, la expresión es siempre falsa, y por ende (n != 1 || n != 0) es siempre verdadera.

}

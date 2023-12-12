package practica

import (
	"fmt"

	"github.com/labora/labora-golang/util"
)

func sortDice() int {
	return util.GenerateRandomNumber(0, 1)
}

func Ej1() {
	// Se necesita generar dos secuencias de números que sigan la siguiente lógica
	// 1. Números en orden crecientes (del 1 en adelante)
	// 2. Números que sean primos (1,2,3,5,7,9,11…)
	// Se pide implementar cada secuencia en una struct que siga cumpla con la interfaz
	// type IntSequence interface {
	// 	Next() int
	// }
	// Por último se pide desarrollar un pequeño programa donde se pueda imprimir los primeros 30 números de una de las dos secuencias según una “tirada de dados”.

	// Tip: Podemos simular una tirada de dado usando rand.Intn(2) (del paquete math/rand).
	fmt.Print("EJERCICIO 1\n\n")
}

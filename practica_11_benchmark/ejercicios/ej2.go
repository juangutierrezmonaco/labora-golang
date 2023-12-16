package practica

import (
	"fmt"
)

func Ej2() {
	// Hemos usado canales para sumar de forma concurrente todos los valores presentes en los elementos de un slice (o array) de números enteros.
	// ¿Qué te parece si ponemos a prueba si la concurrencia en este caso es realmente efectiva? Bueno, es una pregunta pero lo tenés
	// que hacer igual… lo lamento 😁 y me gustaría que discutamos luego sobre CUANDO la concurrencia mejora la performance.
	// Recordar que concurrencia no es lo mismo que paralelismo!.

	// Podés guiarte con este código(https://github.com/vituchon/labora-golang-course/blob/master/meeting-concurrency/slices/slice_test.go) (Ver `func BenchmarkSumUsingChannelI`)
	fmt.Print("EJERCICIO 2\n\n")
}

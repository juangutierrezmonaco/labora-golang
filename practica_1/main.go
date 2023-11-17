package main

import (
	"fmt"
	"github.com/labora/labora-golang/practica_1/ejercicios"
	"github.com/labora/labora-golang/practica_1/util"
)

func validateAndSetOption(inf int, sup int) int {
	var op int
	fmt.Printf("Hola! Ingresá qué ejercicio querés correr (%d a %d, 0 para SALIR): ", inf+1, sup)
	fmt.Scan(&op)
	for op < inf || op > sup {
		fmt.Printf("Opción incorrecta: Ingrese nuavemente (%d a %d, 0 para SALIR): ", inf+1, sup)
		fmt.Scan(&op)
	}
	return op
}

func main() {
	op := validateAndSetOption(0, 7)

	if op != 0 {
		functions := []func(){
			practica.Ej1,
			practica.Ej2,
			practica.Ej3,
			practica.Ej4,
			practica.Ej5,
			practica.Ej6,
			practica.EjGrupo,
		}

		functions[op-1]()

		var runAgain string
		fmt.Print("\n¿Querés correr volver al INICIO? S/s: ")
		fmt.Scan(&runAgain)
		if runAgain == "S" || runAgain == "s" {
			util.ClearConsole()
			main()
		} else {
			fmt.Println("¡Adiós!")
		}
	}

}

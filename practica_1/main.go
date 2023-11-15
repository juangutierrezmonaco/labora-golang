package main

import (
	"fmt"
	"github.com/labora/labora-golang/practica_1/ejercicios"
)

func validateAndSetOption(inf int, sup int) int {
	var op int
	fmt.Printf("Hola! Ingresá qué ejercicio querés correr (%d a %d): ", inf, sup)
	fmt.Scan(&op)
	for op < inf || op > sup {
		fmt.Printf("Opción incorrecta: Ingrese nuavemente (%d a %d): ", inf, sup)
		fmt.Scan(&op)
	}
	return op
}

func main() {
	op := validateAndSetOption(1, 7)
	switch op {
	case 1:
		practica.Ej1()
	case 2:
		practica.Ej2()
	case 3:
		practica.Ej3()
	case 4:
		practica.Ej4()
	case 5:
		practica.Ej5()
	case 6:
		practica.Ej6()
	default:
		practica.EjGrupo()
	}

}

package practica

import (
	"fmt"

	"github.com/labora/labora-golang/util"
)

func findNumberDivisors(num int) []int {
	var divisors []int
	for i := 1; i < num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func Ej1() {
	// Si quisiera hacer una función para obtener los divisores propios de un número natural (ver ejercicio 1 de la práctica funciones) ¿se pueden usar arreglos como
	// tipo de dato de retorno de la función que hace que calcula los divisores propios? Discutan con compañeros.. sin ejercer violencia.
	// “Dime de un hombre que recurre a la violencia y sabré que no es hombre sabio”
	fmt.Print("EJERCICIO 1\n\n")
	
	var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)

	divs := findNumberDivisors(num)
	fmt.Printf("El número %d tiene por divisores: %v\n", num, util.JoinIntArr(divs, ", "))
}

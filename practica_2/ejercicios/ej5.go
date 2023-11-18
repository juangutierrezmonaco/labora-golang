package practica

import (
	"fmt"
)

func calculateTerm(num *float32, lim float32) float32 {
	numBackup := *num
	*num = *num - lim
	if numBackup-lim < 0 && numBackup < 0 {
		return 0
	} else if numBackup < lim {
		return numBackup
	} else {
		return lim
	}
}

func Ej5() {
	// Redactar un algoritmo que pida al usuario el ingreso de una serie de números reales e imprima por pantalla el resultado de sumarlos. La serie finaliza cuando el
	// usuario ingresa el número cero. Comparando el ejercicio anterior, qué analogías y diferencias encuentra. Debata con sus compañeros, 
	// siguiendo las premisas de no violencia. Recuerde, "dime de un hombre sabio y sabré que no es un hombre que recurre a la violencia"
	fmt.Print("EJERCICIO 5\n\n")

	var num float64 = -1
	var acum float64 = 0

	fmt.Println("Ingrese números reales. Para salir presione 0: ")
	for num != 0 {
		fmt.Scan(&num)
		acum += num
	}
	fmt.Printf("La suma da por resultado: %.2f\n", acum)
}

package practica

import (
	"fmt"
	"math"
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
	// 5. Expresar a un número cualquiera como la suma de una serie de números siguiendo las restricciones impuestas a continuación.
	//	Sea X el número.
	//	X = S1 + S2 + S3 + S4 + S5
	//	Donde
	//	0 ≤ S1 ≤ 50
	//	0 ≤ S2 ≤ 50
	//	0 ≤ S3 ≤ 600
	//	0 ≤ S4 ≤ 800
	//	0 ≤ S5 < (Infinito)
	//	De esta forma, si X vale 120, entonces
	//	S1 = 50, S2 = 50, S3 = 20, S4 = 0 y S5 = 0
	//	Si X vale 860, entonces
	//	S1 = 50, S2 = 50, S3 = 600, S4 = 160 y S5 = 0
	fmt.Print("EJERCICIO 5\n\n")

	var s1, s2, s3, s4, s5 float32
	var num float32
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&num)
	fmt.Printf("Los términos del número %.0f son: \n", num)
	s1 = calculateTerm(&num, 50)
	s2 = calculateTerm(&num, 50)
	s3 = calculateTerm(&num, 600)
	s4 = calculateTerm(&num, 800)
	s5 = calculateTerm(&num, float32(math.Inf(1)))
	fmt.Printf("S1 = %.2f, S2 = %.2f, S3 = %.2f, S4 = %.2f y S5 = %.2f\n", s1, s2, s3, s4, s5)
}

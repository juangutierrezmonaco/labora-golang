package practica

import "fmt"

func sum(t1 float32, t2 float32, t3 float32) float32 {
	return t1 + t2 + t3
}

func avg(t1 float32, t2 float32, t3 float32) float32 {
	return sum(t1, t2, t3) / 3
}

func Ej3() {
	// Dados tres valores ambientales de temperatura, desarrollar un algoritmo para calcular e informar la suma y el promedio de dichos valores.
	fmt.Print("EJERCICIO 3\n\n")

	var t1, t2, t3 float32
	fmt.Print("Ingrese un temperatura 1: ")
	fmt.Scan(&t1)
	fmt.Print("Ingrese un temperatura 2: ")
	fmt.Scan(&t2)
	fmt.Print("Ingrese un temperatura 3: ")
	fmt.Scan(&t3)

	fmt.Printf("La suma de las temperaturas %.2f, %.2f y %.2f es %.2f.\n", t1, t2, t3, sum(t1, t2, t3))
	fmt.Printf("El promedio de las temperaturas %.2f, %.2f y %.2f es %.2f.\n", t1, t2, t3, avg(t1, t2, t3))
}

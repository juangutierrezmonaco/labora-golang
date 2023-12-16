package main

import (
	"fmt"
)

func main() {
	// Se desea poner a prueba a los algoritmos de rotación de arrays (o slices). Para esto se pide escribir test unitarios que los pongan a prueba.
	// Podés guiarte con este código:
	fmt.Print("EJERCICIO 3\n\n")
	slice := []int{1, 2, 3, 4, 5}
	rotateAllElementsToRight(&slice)
	fmt.Println(slice)
}

func rotateAllElementsToRight[T any](slice *[]T) {
	lastPos := len(*slice) - 1
	*slice = append((*slice)[lastPos:], (*slice)[:lastPos]...)
}

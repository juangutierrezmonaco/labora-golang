package practica

import "fmt"

func rectangleArea(length float32, width float32) float32 {
	return length * width
}

func rectanglePerimeter(length float32, width float32) float32{
	return 2 * (length * width)
}

func EjGrupo() {
	// Escribe un programa que calcule el área y el perímetro de un rectángulo. El programa debe pedir al usuario que introduzca la longitud y la anchura 
	// del rectángulo y  luego calcular y mostrar el área y el perímetro.
	// Especificaciones:
	// Solicita al usuario que ingrese la longitud y la anchura del rectángulo.
	// Calcula el área del rectángulo (longitud * anchura).
	// Calcula el perímetro del rectángulo (2 * (longitud + anchura)).
	// Muestra los resultados (área y perímetro) al usuario.
	// Ejemplo de Salida:
	// Ingrese la longitud del rectángulo: 5 
	// Ingrese la anchura del rectángulo: 3 
	// 
	// 
	// El área del rectángulo es: 15 
	// El perímetro del rectángulo es: 16
	fmt.Print("EJERCICIO 7\n\n")

	var length, width float32
	fmt.Print("Ingrese la longitud del rectángulo: ")
	fmt.Scan(&length)
	fmt.Print("Ingrese la anchura del rectángulo: ")
	fmt.Scan(&width)

	fmt.Printf("El perímetro del rectángulo de longitud %.2f y ancho %.2f es: %.2f\n", length, width, rectanglePerimeter(length, width))
	fmt.Printf("El área del rectángulo de longitud %.2f y ancho %.2f es: %.2f\n", length, width, rectangleArea(length, width))
}

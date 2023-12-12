package practica

import (
	"fmt"
)

type PersonA struct {
	Name  string
	Age   int
	City  string
	Phone string
}

func (p *PersonA) changeCity(newCity string) {
	if newCity != p.City {
		p.City = newCity
	}
}

func Ej1() {
	// Define una estructura llamada Persona con los siguientes campos: nombre, edad, ciudad y telefono.
	// Crea dos variables de tipo Persona con valores iniciales distintos.
	// Imprime los valores de ambas variables por pantalla.

	// Define una función llamada cambiarCiudad para cambiar la ciudad de la persona. Si la ciudad es la misma, no hace nada.
	// Llama a la función cambiarCiudad con una de las personas creadas anteriormente y una ciudad distinta a la actual.
	// Imprime los valores de ambas variables por pantalla para comprobar que se ha actualizado el campo ciudad de la persona correspondiente si la ciudad es distinta,
	// o que no se ha actualizado si la ciudad es la misma.

	// Aquí está la salida del programa para los valores de entrada dados:
	// Persona 1: {Juan 30 Madrid 555-1234}
	// Persona 2: {Ana 25 Barcelona 555-5678}

	// Persona 1 con ciudad actualizada: {Juan 30 Valencia 555-1234}
	// Persona 2: {Ana 25 Barcelona 555-5678}
	fmt.Print("EJERCICIO 1\n\n")

	person1 := PersonA{"Juan", 30, "Madrid", "555-1234"}
	person2 := PersonA{"Ana", 25, "Barcelona", "555-5678"}

	fmt.Printf("Persona 1: %v\n", person1)
	fmt.Printf("Persona 2: %v\n\n", person2)

	person1.changeCity("Valencia")

	fmt.Printf("Persona 1 con ciudad actualizada: %v\n", person1)
	fmt.Printf("Persona 2: %v\n", person2)

}

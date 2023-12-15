package practica

import (
	"fmt"

	"github.com/labora/labora-golang/util"
)

func validateNonEmptyOrLengthyString(str string) error{
	switch {
	case len(str) > 100:
		return fmt.Errorf("%s\n", "El texto es muy largo.")
	case str == "":
		return fmt.Errorf("%s\n", "El texto está vacío.")
	}
	return nil
}

func Ej2() {
	// Crea una función que acepte un string y devuelva un error si el string está vacío o si es demasiado largo (por ejemplo, más de 100 caracteres).
	fmt.Print("EJERCICIO 2\n\n")

	fmt.Print("Ingrese un texto: ")
	str := util.ScanSentence()

	err := validateNonEmptyOrLengthyString(str)
	if (err != nil) {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("El texto \"%s\" no está vacío ni es mas largo que 100 caracteres.\n", str)
	}
}

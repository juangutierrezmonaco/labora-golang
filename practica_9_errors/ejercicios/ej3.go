package practica

import (
	"fmt"

	"github.com/labora/labora-golang/util"
)

func strToInt(str string) (int, error) {
	res, err := util.Atoi(str)
	if err != nil {
		err = fmt.Errorf("El texto ingresado \"%s\" no es un número entero.", str)
	}
	return res, err
}

func Ej3() {
	// Escribe una función que convierta un string a un número entero (int). La función debe devolver un error si el string
	// no puede convertirse a un número. Prueba tu función con diferentes strings, incluyendo algunos que no sean números.
	// Usa panic y recover para manejar los errores
	fmt.Print("EJERCICIO 3\n\n")

	defer recoverAndPrintError()

	fmt.Print("Ingrese un texto para convertirlo a número entero: ")
	str := util.ScanSentence()
	num, err := strToInt(str)
	if err != nil {
		panic(err)
	}

	fmt.Printf("El número es: %d\n", num)
}

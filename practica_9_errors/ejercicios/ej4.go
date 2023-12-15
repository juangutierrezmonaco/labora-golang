package practica

import (
	"fmt"
	"os"
)

func openOrCreateFile(fileName string, path string) *os.File {
	fullPath := fmt.Sprintf("%s/%s", path, fileName)
	// Opening file
	file, err := os.Open(fullPath)
	if err != nil {
		// Creating if it doesn't exists
		file, err = os.Create(fullPath)
		if err != nil {
			panic(fmt.Sprintf("No se pudo crear el archivo \"%s\" en la ruta \"%s\" (info: %s)", fileName, path, err.Error()))
		}
		fmt.Printf("Se creó el archivo: %s en la ruta %s\n", fileName, path)
	}
	return file
}

func Ej4() {
	// Escribe un programa que intente abrir un archivo cuyo nombre y directorio se pasa como argumento en la línea de comandos. Si el archivo no existe,
	// el programa debe crearlo. Si hay algún otro error al abrir el archivo, el programa debe informar del error. No olvides cerrar el archivo
	// correctamente en todos los casos.
	fmt.Print("EJERCICIO 4\n\n")

	defer recoverAndPrintError()

	args := os.Args
	// go run . [arg1=name] [arg2=path]
	// The first arg is the path of the build exe
	if len(args) != 3 {
		panic("Debe ingresar el nombre del archivo y la ruta donde se encuentra como argumento en la línea de comandos")
	}

	fileName := os.Args[1]
	path := os.Args[2]

	file := openOrCreateFile(fileName, path)
	defer file.Close()

	fmt.Println("El archivo se abrió correctamente.")
}

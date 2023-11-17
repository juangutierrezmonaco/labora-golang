package main

import (
	"github.com/labora/labora-golang/practica_1/ejercicios"
	"github.com/labora/labora-golang/util"
)

func main() {
	functions := []func(){
		practica.Ej1,
		practica.Ej2,
		practica.Ej3,
		practica.Ej4,
		practica.Ej5,
		practica.Ej6,
		practica.EjGrupo,
	}

	util.RunMenu(0, 7, functions)
}

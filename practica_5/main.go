package main

import (
	"github.com/labora/labora-golang/practica_5/ejercicios"
	"github.com/labora/labora-golang/util"
)

/**********************************************************************************/
/*                               PRÁCTICA DE STRUCTS                              */
/**********************************************************************************/

func main() {

	functions := []func(){
		practica.Ej1,
		practica.Ej2,
		practica.EjPairProgramming,
	}

	util.RunMenu(0, len(functions), functions)
}

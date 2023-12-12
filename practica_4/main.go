package main

import (
	"github.com/labora/labora-golang/practica_4/ejercicios"
	"github.com/labora/labora-golang/util"
)

/**********************************************************************************/
/*                                PRÁCTICA DE ARRAYS                              */
/**********************************************************************************/

func main() {

	functions := []func(){
		practica.Ej1,
		practica.Ej2,
		practica.Ej3,
		practica.Ej4,
		practica.EjPairProgramming,
	}

	util.RunMenu(0, len(functions), functions)
}

package main

import (
	"github.com/labora/labora-golang/practica_7_interfaces/ejercicios"
	"github.com/labora/labora-golang/util"
)

/**********************************************************************************/
/*                            PRÁCTICA DE INTERFACES                              */
/**********************************************************************************/

func main() {

	functions := []func(){
		practica.Ej1_2_3,
		practica.Ej4,
		practica.Ej5,
		practica.EjPairProgramming,
	}

	util.RunMenu(0, len(functions), functions)
}

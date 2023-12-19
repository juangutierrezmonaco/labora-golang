package main

import (
	"github.com/labora/labora-golang/practica_12_race_conditions/ejercicios"
	"github.com/labora/labora-golang/util"
)

/**********************************************************************************/
/*                         PRÁCTICA DE RACE CONDITIONS                            */
/**********************************************************************************/

func main() {
	functions := []func(){
		practica.Ej1,
		practica.Ej2,
		practica.Ej3,
		practica.Ej3b,
	}

	util.RunMenu(0, len(functions), functions)
}

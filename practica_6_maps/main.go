package main

import (
	"github.com/labora/labora-golang/practica_6_maps/ejercicios"
	"github.com/labora/labora-golang/util"
)

/**********************************************************************************/
/*                               PRÁCTICA DE MAPAS                                */
/**********************************************************************************/

func main() {

	functions := []func(){
		practica.Ej1,
	}

	util.RunMenu(0, len(functions), functions)
}

package main

import (
	"github.com/labora/labora-golang/practica_8_concurrency/ejercicios"
	"github.com/labora/labora-golang/util"
)

/**********************************************************************************/
/*                          PRÁCTICA DE CONCURRENCIA                              */
/**********************************************************************************/

func main() {

	functions := []func(){
		practica.Ej1,
		practica.Ej2,
		practica.Ej3,
	}

	util.RunMenu(0, len(functions), functions)
}

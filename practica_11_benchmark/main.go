package main

import (
	"github.com/labora/labora-golang/practica_11_benchmark/ejercicios"
	"github.com/labora/labora-golang/util"
)

/**********************************************************************************/
/*                           PRÁCTICA DE BENCHMARK                                */
/**********************************************************************************/

func main() {

	functions := []func(){
		practica.Ej1,
		practica.Ej2,
	}

	util.RunMenu(0, len(functions), functions)
}

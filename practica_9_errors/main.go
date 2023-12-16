package main

import (
	"github.com/labora/labora-golang/practica_9_errors/ejercicios"
	"github.com/labora/labora-golang/util"
)

/**********************************************************************************/
/*                             PRÁCTICA DE ERRORES                                */
/**********************************************************************************/

func main() {

	functions := []func(){
		practica.Ej1,
		practica.Ej2,
		practica.Ej3,
		practica.Ej4,
		practica.Ej5,
	}

	util.RunMenu(0, len(functions), functions)
}

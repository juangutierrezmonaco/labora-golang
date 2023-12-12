package main

import (
	"github.com/labora/labora-golang/practica_3/ejercicios"
	"github.com/labora/labora-golang/util"
)

/**********************************************************************************/
/*                                PRÁCTICA DE FUNCIONES                           */
/**********************************************************************************/

func main() {

	functions := []func(){
		practica.Ej1,
		practica.Ej2,
		practica.Ej3,
		practica.Ej4,
		practica.Ej5,
		practica.Ej6,
		practica.Ej7,
		practica.Ej8,
		practica.Ej9,
		practica.Ej10,
		practica.Ej11,
	}

	util.RunMenu(0, len(functions), functions)
}

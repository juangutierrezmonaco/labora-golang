package main

import (
	"github.com/labora/labora-golang/practica_1/ejercicios"
	"github.com/labora/labora-golang/util"
)

/**********************************************************************************/
/*                      PRÁCTICA DE SENTENCIAS CONDICIONALES                      */
/**********************************************************************************/

func main() {
	functions := []func(){
		practica.Ej1,
		practica.Ej2,
		practica.Ej3,
		practica.Ej4,
		practica.Ej5,
		practica.Ej6,
		practica.EjPairProgramming,
	}

	util.RunMenu(0, len(functions), functions)
}

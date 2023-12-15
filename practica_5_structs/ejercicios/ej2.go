package practica

import (
	"fmt"
	"github.com/labora/labora-golang/util"
)

type IntegerSet struct {
	Numbers []int
	Length  int
}

func (intSet *IntegerSet) Add(value int) {
	intSet.Numbers = append(intSet.Numbers, value)
	intSet.Length++
}

func (intSet *IntegerSet) RemoveAt(index int) bool {
	isInBoundaries := intSet.isInBoundaries(index)
	if isInBoundaries {
		intSet.Numbers = append(intSet.Numbers[:index], intSet.Numbers[index+1:]...)
		intSet.Length--
	}
	return isInBoundaries
}

func (intSet *IntegerSet) Lenght() int {
	return intSet.Length
}

func (intSet *IntegerSet) Set(value int, index int) bool {
	isInBoundaries := intSet.isInBoundaries(index)
	if isInBoundaries {
		intSet.Numbers[index] = value
	}
	return isInBoundaries
}

func (intSet *IntegerSet) ToString() string {
	return util.JoinIntArr(intSet.Numbers, ", ")
}

func (intSet *IntegerSet) isInBoundaries(index int) bool {
	return index >= 0 && index < intSet.Length
}

func Ej2() {
	// Defina una estructura de datos para manejar conjuntos de enteros de la siguiente forma.
	// Add(value int) int: Agrega al final. Se incrementa en uno la longitud.
	// RemoveAt(index int) bool: True si lo borró, false si no pudo (no importa el motivo). Si se borra entonces se decrementa en uno la longitud
	// Lenght() int: Cantidad de elementos que tiene. Esta determinada por la cantidad de valores que se agregan menos aquellos que se borran.
	// Set(value int, index int) bool: True si lo modificó, false si no pudo (no importa el motivo)
	// ToString() string: Libre elección. La función strings.Join (del paquete strings) puede ser de útil.
	// Para pensar: Tanto el RemoveAt y el Set son métodos cuya ejecución puede “fallar” si se pasa como index un valor que esté fuera de los “límites”
	// (dado por la longitud)… Yo pregunto entonces 🤔 ¿Es esto algo entonces que pueda ser codificado en una función aparte de forma tal que
	// se re utilice en ambos métodos? ==> isInBoundaries(index) bool
	fmt.Print("EJERCICIO 2\n\n")

	intSet := IntegerSet{}

	fmt.Println("Agregando el elemento '1' y '2'")
	intSet.Add(1)
	intSet.Add(2)
	fmt.Printf("Arreglo: %s\n", intSet.ToString())

	fmt.Println("Eliminado el elemento de la posición '1'")
	intSet.RemoveAt(1)
	fmt.Printf("Arreglo: %s\n", intSet.ToString())

	fmt.Println("Eliminado el elemento de la posición '10' (posición fuera de rango)")
	intSet.RemoveAt(1)
	fmt.Printf("Arreglo: %s\n", intSet.ToString())

	fmt.Println("Mostrando la longitud del arreglo")
	fmt.Printf("Longitud: %d\n", intSet.Length)

	fmt.Println("Setteando el elemento de la posición '0' con un 20")
	intSet.Set(20, 0)
	fmt.Printf("Arreglo: %s\n", intSet.ToString())
}

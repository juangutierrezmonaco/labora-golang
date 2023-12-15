package practica

import (
	"fmt"
)

func removeValueInPos(arr []int, pos int){
	for i := pos; i < len(arr) - 1; i++ {
		arr[i] = arr[i+1] 
	}
	lastPos := len(arr) - 1
	arr[lastPos] = 0
}

func Ej2() {
	// Haga que reciba un arreglo y una posición y “borre” el valor que hay en dicha posición. Por “borrar” entiendase que no se quita el elemento sino se mueven todos 
	// los elementos que están a la derecha un pasito a la izquierda, rellenado con el valor por defecto el extremo derecho. O sea…
  // Si tengo el arreglo int[5]{10,20,30,40,50} y “borro” el elemento de la posición 2 (recordar que se indexan desde 0) quedaría ⇒ int[5]={10,20,40,50,0}
	fmt.Print("EJERCICIO 2\n\n")
	
	var pos int
	arr := []int{10, 20, 30, 40, 50}

	fmt.Printf("Ingrese la posición a eliminar del arreglo (%v): ", arr)
	fmt.Scan(&pos)
	removeValueInPos(arr, pos)
	fmt.Println(arr)
}

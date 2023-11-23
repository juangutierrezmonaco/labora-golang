package practica

import (
	"fmt"
)

func toInt(cad string) int {
	long := len(cad)
	i, pot, res := 0, 1, 0

	for i < long {
		c := cad[long-i-1]
		res += int(c-'0') * pot
		pot *= 10
		i++
	}

	return res
}

func Ej5() {
	// Comenta que realiza el siguiente programa. Como se pide entrada de datos por teclado, no va a funcionar en playground remoto,
	// por eso pido que se copien el código a un archivo en su computadora y corran el programa con el comando go run <archivo.go>
	fmt.Print("EJERCICIO 5\n\n")

	var cadena string
	fmt.Print("Ingrese un número como cadena: ")
	fmt.Scan(&cadena)

	resultado := toInt(cadena)
	fmt.Printf("El resultado como entero es: %d\n", resultado)

	// Pasa un número en string (ENTERO) al tipo int. Esto lo hace haciendo lo siguiente:
	// 	- 1) Recorre cada uno de los dígitos del número (en string) de la cadena de forma regresiva (es decir, del último dígito al primero)
	// 	- 2) Agarra ese dígito, lo transforma a número entero haciendo el cast restándole '0' (48 en ASCII) para evitar pasarlo a su representación ASCII.
	// 	- 3) Lo múltiplica por la potencia de 10 correspondiente a la posición, siendo esta: 1, 10, 100, 1000, ...
	// 	- 4) Lo suma al resultado (inicialmente 0), y como se va sumando por potencia de diez queda completo el número.
	// 	- 5) En cada ciclo, se aumenta el contador en 1 y la potencia en 10.

	// 	** Ejemplo: "1234" **
	// 	- "1234" ==> '1', '2', '3', '4'
	// 	- '4' ==> 4 * 1			=    4
	// 	- '3' ==> 3 * 10		=   30
	// 	- '2' ==> 2 * 100		=  200
	// 	- '1' ==> 1 * 1000	= 1000
	// 	- 							res	= 1234

}

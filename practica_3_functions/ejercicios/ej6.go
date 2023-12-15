package practica

import (
	"fmt"
)

// Devuelve
func unAlgoritmo(a int) int {
	// Duplica el número
	quieroCunfundir := a * 2

	for a != 0 {
		// Hace un loop "a" veces, pues el decremento viene después de entrar en el ciclo
		a--
		// Acumula en la variable que ya se duplicó previamente, la suma del doble de los números naturales hasta 0 partiendo de a - 1
		//  ==> (a * 2) + (a - 1) * 2 + (a - 2) * 2 + ... + (a - a) * 2		=		2a + 2*Σ(1, a - 1)		= 	2 * (a + Σ(1, a - 1))
		quieroCunfundir += a * 2
	}
	// Aplicando suma de Gauss: Σ(1, a - 1) = (a-1)*(a-1+1) / 2 = a*(a-1) / 2

	// Resultando en: 2 * (a + Σ(1, a - 1)) = 2a + 2*(a*(a-1) / 2	) = 2a + a*(a-1) = 2a + a^2 - a = a^2 + a = a * (a + 1)

	// Esta función responde a la función matemática: f(a) = a * (a + 1) --> Devuelve el número multiplicado por el siguiente.

	return quieroCunfundir
}

// Equivalente a unAlgoritmo
func multiplyNumberByNextNumber(num int) int {
	res := num * 2
	for num != 0 {
		num--
		res += num * 2
	}
	return res
}

func unaCosa(valor int) {
	resultado := unAlgoritmo(valor) * unAlgoritmo(valor)
	fmt.Println("El resultado es:", resultado)
}

// Equivalente a unaCosa
func multiplySquareOfNumberBySquareOfNextNumber(valor int) {
	resultado := multiplyNumberByNextNumber(valor) * multiplyNumberByNextNumber(valor)
	fmt.Println("El resultado es:", resultado)
}

func Ej6() {
	// Comenta que realiza el siguiente algoritmo.
	// Una vez entendido lo que hace el algoritmo proponga nuevos nombres para los identificadores que describan mejor su uso.
	fmt.Print("EJERCICIO 6\n\n")

	// Devuelve el cuadrado del número * el cuadrado del siguiente número:
	// EJ: num = 7 --> res = 7^2*8^2
	multiplySquareOfNumberBySquareOfNextNumber(2)
}

package practica

import (
	"fmt"

	"github.com/labora/labora-golang/practica_3/ejercicios"
	"github.com/labora/labora-golang/util"
)

func rollATriangularDice() int {
	return util.GenerateRandomNumber(0, 2)
}

type IntSequence interface {
	Next() int
	Title() string
	Fulfill(n int) bool
}

type LinealIntSequence struct {
	last int
}

func (s *LinealIntSequence) Next() int {
	s.last++
	return s.last
}

func (s *LinealIntSequence) Title() string {
	return "Secuencia lineal de números enteros."
}

func (s *LinealIntSequence) Fulfill(n int) bool {
	return true
}

type PrimeSequence struct {
	last int
}

func (s *PrimeSequence) Next() int {
	newPrimeFounded := false
	for i := s.last + 1; !newPrimeFounded; i++ {
		if s.Fulfill(i) {
			newPrimeFounded = true
			s.last = i
		}
	}
	return s.last
}

func (s *PrimeSequence) Title() string {
	return "Secuencia de números primos."
}

func (s *PrimeSequence) Fulfill(n int) bool {
	return practica.IsPrimeNumber(n) || n == 1
}

type EvenSequence struct {
	last int
}

func (s *EvenSequence) Next() int {
	s.last += 2
	return s.last
}

func (s *EvenSequence) Title() string {
	return "Secuencia de números pares."
}

func (s *EvenSequence) Fulfill(n int) bool {
	return n%2 == 0
}

func Ej1_2_3() {
	// Se necesita generar dos secuencias de números que sigan la siguiente lógica
	// 1. Números en orden crecientes (del 1 en adelante)
	// 2. Números que sean primos (1,2,3,5,7,9,11…)
	// Se pide implementar cada secuencia en una struct que siga cumpla con la interfaz
	// type IntSequence interface {
	// 	Next() int
	// }
	// Por último se pide desarrollar un pequeño programa donde se pueda imprimir los primeros 30 números de una de las dos secuencias según una “tirada de dados”.

	// Tip: Podemos simular una tirada de dado usando rand.Intn(2) (del paquete math/rand).

	// Ahora se desea que se imprima un mensaje que diga cuál de las dos secuencias se tomó (según la tirada de dados).
	// 💡 TIP: Se le puede declarar un método `Title() string` en la interface `IntSequence` y hacer que ambas secuencias lo implementen.

	// De forma más general se desea tener una secuencia que me dé números que cumplan una condición (predicado).
	// Dicha condición será implementada en estructuras (o cualquier otro tipo de dato)  que tengan el un método con la siguiente firma `Fulfill(n int) bool`.
	// Fulfill es un término en inglés que comúnmente se emplea para afirmar que algo cumple una condición.
	// Además de tener tipos de datos que tengan una función para implementar condiciones de:
	// Números cualquier sea (para secuencias de enteros lineales)
	// Números primos
	// Ahora quiero que agreguen un nuevo tipo que tenga una función que implemente la condición de
	// Números pares (2,4,6…)
	fmt.Print("EJERCICIO 1, 2 y 3\n\n")

	intSequences := []IntSequence{&LinealIntSequence{}, &PrimeSequence{}, &EvenSequence{}}
	selectedSequence := intSequences[rollATriangularDice()]

	fmt.Println(selectedSequence.Title())
	fmt.Print("[ ")
	for i := 0; i < 30; i++ {
		fmt.Printf("%d ", selectedSequence.Next())
	}
	fmt.Print("]")
}

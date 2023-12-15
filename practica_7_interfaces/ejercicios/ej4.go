package practica

import (
	"fmt"

	practica "github.com/labora/labora-golang/practica_3_functions/ejercicios"
)

type NumberPredicate interface {
	Fulfill(n int) bool
}

type IntSequenceB interface {
	Next() int
	Title() string
}

type Sequence struct {
	current         int
	numberPredicate NumberPredicate
	typeName        string
	direction       int // 0 asc, 1 desc
}

func (s *Sequence) Next() int {
	var next int
	fulfillsPredicate := false
	var n int
	if s.direction == 0 {
		n = s.current + 1
	} else {
		n = s.current - 1
	}
	for !fulfillsPredicate {
		if s.numberPredicate.Fulfill(n) {
			next = n
			fulfillsPredicate = true
		}
		if s.direction == 0 {
			n++
		} else {
			n--
		}
	}
	s.current = next
	return s.current
}

func (s *Sequence) Title() string {
	title := "Soy una secuencia del tipo: " + s.typeName
	if s.direction == 0 {
		title += "(Asc)."
	} else {
		title += "(Desc)."
	}
	return title
}

type Primes struct{}

func (_ *Primes) Fulfill(n int) bool {
	return practica.IsPrimeNumber(n) || n == 1
}

type Integers struct{}

func (_ *Integers) Fulfill(n int) bool {
	return true
}

type Evens struct{}

func (_ *Evens) Fulfill(n int) bool {
	return n%2 == 0
}

func Ej4() {
	// ¿Qué pasaría si quiero una secuencia pero en el orden inverso? 🤔. ¿Acaso es mucho pedir…? Creo que estoy teniendo sed de combinar dos capacidades:
	// por un lado el orden en el que se generan los números: creciente o decreciente y, por otro lado, un predicado que me diga si un número cumple con tal condición.
	// Se desea implementar una estructura Sequence que tenga como campos un generador y un predicado y que tenga un método llamado Next()
	// que devuelva el próximo número que se genera usando el generador y que cumpla con el predicado.
	fmt.Print("EJERCICIO 4\n\n")

	direction := 1
	current := 100

	intSequences := []IntSequenceB{
		&Sequence{numberPredicate: &Integers{}, typeName: "Números enteros", direction: direction, current: current},
		&Sequence{numberPredicate: &Primes{}, typeName: "Números primos", direction: direction, current: current},
		&Sequence{numberPredicate: &Evens{}, typeName: "Números pares", direction: direction, current: current},
	}
	selectedSequence := intSequences[rollATriangularDice()]

	fmt.Println(selectedSequence.Title())
	fmt.Print("[ ")
	for i := 0; i < 30; i++ {
		fmt.Printf("%d ", selectedSequence.Next())
	}
	fmt.Print("]")

}

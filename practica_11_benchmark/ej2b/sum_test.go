package main

import (
	"fmt"
	"math/rand"
	"testing"
)

type Test struct {
	numberOfGoroutines int
}

func generateTests() ([]int, []Test) {
	slice := rand.Perm(10000000)
	tests := []Test{}
	for i := 1; i <= len(slice); i *= 10 {
		t := Test{
			numberOfGoroutines: i,
		}
		tests = append(tests, t)
	}
	return slice, tests
}

// para correr los benchmark
// $slices/go test -benchmem -bench . github.com/vituchon/labora-golang-course/meeting-concurrency/slices
// $slices/go test -bench .  ./slices // todos los benchs
// $slices/go test -bench=BenchmarkSumUsingChannel ./slices // los que sigan el patron
// se puede agregar flag -run=none (para que no ejecute tests!) => slices/go test -run=none  -bench .  ./slices
// https://apuntes.de/golang/pruebas-benchmark/#gsc.tab=0
// https://www.golinuxcloud.com/golang-benchmark/

func BenchmarkSumWithNGoroutines(b *testing.B) {
	values, tests := generateTests()
	len := len(values)
	// Primero vemos cuánto tarda sin gorutinas
	b.Run(fmt.Sprintf("Sumando sin gorutinas %d valores: ", len), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sumValues(&values)
		}
	})

	// Ahora testeamos cuánto demora si dividimos en N gorutinas
	for _, test := range tests {
		b.Run(fmt.Sprintf("Sumando a través de %d gorutinas %d valores: ", test.numberOfGoroutines, len), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SumValuesWithNGoroutines(&values, test.numberOfGoroutines)
			}
		})
	}
}


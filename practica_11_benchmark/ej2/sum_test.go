package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/labora/labora-golang/util"
)

func TestSumsWork(t *testing.T) {
	fmt.Printf("Corriendo el test: TestSumsWork...\n")

	type Test struct {
		values   []int
		expected int
	}

	tests := []Test{
		{
			values:   []int{1, 2, 3},
			expected: 6,
		},
		{
			values:   []int{1, 2, 3, 4},
			expected: 10,
		},
		{
			values:   []int{3, 2, 1, 5, 7, 8, 9},
			expected: 35,
		},
		{
			values:   []int{37, 28, -1, 10, -8, 0},
			expected: 66,
		},
	}
	for _, test := range tests {
		title := util.JoinIntArr(test.values, ", ")
		t.Logf("\n=====\tRunning unit test: [%s]\t=====\n", title)

		generated := SumWithoutUsingChannel(&test.values)
		if test.expected != generated {
			t.Errorf("\nGenerated: '%+v' vs Expected: '%+v'", generated, test.expected)
		}

		generated = SumUsingChannel(&test.values)
		if test.expected != generated {
			t.Errorf("\nGenerated: '%+v' vs Expected: '%+v'", generated, test.expected)
		}

	}
}

// para correr los benchmark
// $slices/go test -benchmem -bench . github.com/vituchon/labora-golang-course/meeting-concurrency/slices
// $slices/go test -bench .  ./slices // todos los benchs
// $slices/go test -bench=BenchmarkSumUsingChannel ./slices // los que sigan el patron
// se puede agregar flag -run=none (para que no ejecute tests!) => slices/go test -run=none  -bench .  ./slices
// https://apuntes.de/golang/pruebas-benchmark/#gsc.tab=0
// https://www.golinuxcloud.com/golang-benchmark/

var slice = rand.Perm(100000000)

func BenchmarkSumUsingChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumUsingChannel(&slice)
	}
}

func BenchmarkSumWithoutUsingChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumWithoutUsingChannel(&slice)
	}
}

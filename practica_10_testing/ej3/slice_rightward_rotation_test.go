package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type Test[T any] struct {
	input    []T
	expected []T
}

func TestRightwardRotationsWorksOk(t *testing.T) {
	fmt.Printf("Corriendo el test: TestRightwardRotationsWorksOk...\n")

	testsWithInts := []Test[int]{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 1, 2, 3, 4},
		},
		{
			input:    []int{1, 2, 3},
			expected: []int{3, 1, 2},
		},
	}

	testsWithFloats := []Test[float64]{
		{
			input:    []float64{1.2, 2, -1},
			expected: []float64{-1, 1.2, 2},
		},
	}

	testsWithStrings := []Test[string]{
		{
			input:    strings.Split("HOLA", ""),
			expected: strings.Split("AHOL", ""),
		},
	}

	runTests(testsWithInts, t)
	runTests(testsWithFloats, t)
	runTests(testsWithStrings, t)
}

func runTests[T any](tests []Test[T], t *testing.T) {
	for _, test := range tests {
		computedSlice := test.input
		rotateAllElementsToRight(&computedSlice)

		if !reflect.DeepEqual(computedSlice, test.expected) {
			t.Errorf("Arreglo esperado(%v) difiere de arreglo computado(%v). Siendo del tipo: %v\n", test.expected, computedSlice, reflect.TypeOf(computedSlice))
		}
	}
}

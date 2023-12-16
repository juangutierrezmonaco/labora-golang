package main

import (
	"fmt"
	"testing"
)

func TestAddNoteWorksOk(t *testing.T) {
	fmt.Printf("Corriendo el test: TestAddNoteWorksOk...\n")

	type Test struct {
		n           int
		expectedSum int
	}

	var tests []Test = []Test{
		{
			n:           3,
			expectedSum: 9,
		},
		{
			n:           -3,
			expectedSum: 0,
		},
		{
			n:           0,
			expectedSum: 0,
		},
	}

	for _, test := range tests {
		computedSum, err := SumNOddNumbers(test.n)
		if computedSum != test.expectedSum {
			t.Errorf("Suma esperada(%d) difiere de suma computada(%d).\n", test.expectedSum, computedSum)
			t.Errorf("El error fue: %s.\n", err)
		}
	}
}

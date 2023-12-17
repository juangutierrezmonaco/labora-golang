package main

import (
	"fmt"
	"testing"
)

func TestSumOddsMethodsDoTheSame(t *testing.T) {
	fmt.Printf("Corriendo el test: TestSumOddsMethodsDoTheSame...\n")

	type Test struct {
		n int
	}

	var tests []Test = []Test{
		{
			n: 0,
		},
		{
			n: 1,
		},
		{
			n: 7,
		},
		{
			n: 21,
		},
		{
			n: 55,
		},
	}
	for _, test := range tests {
		val1, _ := SumNOddNumbersLoop1(test.n)
		val2, _ := SumNOddNumbersLoop2(test.n)
		val3, _ := SumNOddNumbersGauss(test.n)
		if !(val1 == val2 && val2 == val3) {
			t.Errorf("Test with %d does not return all same values, they were v1=%d,  v2=%d,  v3=%d", test.n, val1, val2, val3)
		}
	}
}

func Benchmark_sumNOddNumbersLoop1_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumNOddNumbersLoop1(100)
	}
}

func Benchmark_sumNOddNumbersLoop2_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumNOddNumbersLoop2(100)
	}
}

func Benchmark_sumNOddNumbersGauss_100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumNOddNumbersGauss(100)
	}
}

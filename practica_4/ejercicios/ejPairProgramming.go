package practica

import (
	"fmt"
	"strings"
)

type DnaChain [][]string

func NewDnaChain() DnaChain {
	dnaChain := make([][]string, 3)
	return dnaChain
}

type Coord struct {
	x int
	y int
}

func clockwiseTraversalIndex(row int, col int) int {
	// Example usage
	indexByIndexes := map[Coord]int{
		{0, 0}: 0,
		{0, 1}: 1,
		{0, 2}: 2,
		{1, 2}: 3,
		{2, 2}: 4,
		{2, 1}: 5,
		{2, 0}: 6,
		{1, 0}: 7,
		{1, 1}: 8,
	}

	return indexByIndexes[Coord{x: row, y: col}]
}

func dnaChainToStr(dnaChain DnaChain) string {
	dnaChainArr := make([]string, 8)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			arrIndex := clockwiseTraversalIndex(i, j)
			if arrIndex != 8 {
				dnaChainArr[arrIndex] = dnaChain[i][j]
			}
		}
	}
	return strings.Join(dnaChainArr, "")
}

func dnaChainsAreEqual(d1 DnaChain, d2 DnaChain) bool {
	// Pass to string the dna chains
	d1Str := dnaChainToStr(d1)
	d2Str := dnaChainToStr(d2)

	// add to the first dna chain, the same dna chain to complete the circular nature of the dna string
	d1Str += d1Str

	// If the first dna chain contains the second, they're equal
	return strings.Contains(d1Str, d2Str)
}

func EjPairProgramming() {
	//Una cadena de ADN se representa como una secuencia circular de bases (adenina, timina, citosina y guanina) que es única para cada ser vivo, por ejemplo:

	// |  A  |  T  |  T  |
	// | --- | --- | --- |
	// |  T  |     |  C  |	=====> [A T T C G T A T]
	// | --- | --- | --- |
	// |  A  |  T  |  G  |

	// Dicha cadena se puede representar como un arreglo de caracteres recorriendola en sentido horario desde la parte superior izquierda:
	// A T G C G T A T
	// Se pide diseñar una clase que represente una secuencia de ADN e incluya un método booleano que nos devuelva **true** si dos cadenas de ADN coinciden.

	// MUY IMPORTANTE: La secuencia de ADN es cíclica, por lo que puede comenzar en cualquier posición. Por ejemplo, las dos secuencias siguientes coinciden:
	// **A** T G C G T A T
	// A T **A** T G C G T

	// Pregunta existencial: ¿la cantidad de combinaciones de ADN debe ser finita? ¿Cuántas combinaciones diferentes puede haber?
	// ¿Puede ocurrir algún día que se empiezan a repetir??
	fmt.Print("EJERCICIO 5\n\n")
	dnaChain1 := DnaChain{
		{"A", "T", "T"},
		{"T", " ", "C"},
		{"A", "T", "G"},
	}
	dnaChain2 := DnaChain{
		{"A", "T", "T"},
		{"T", " ", "C"},
		{"A", "T", "G"},
	}
	dnaChain3 := DnaChain{
		{"A", "T", "T"},
		{"F", " ", "C"},
		{"A", "T", "G"},
	}
	dnaChain4 := DnaChain{
		{"G", "T", "A"},
		{"C", " ", "T"},
		{"T", "T", "A"},
	}

	fmt.Println(dnaChainsAreEqual(dnaChain1, dnaChain2)) // true, they're the same
	fmt.Println(dnaChainsAreEqual(dnaChain1, dnaChain3)) // false, they're different
	fmt.Println(dnaChainsAreEqual(dnaChain1, dnaChain4)) // true, they're the same cyclically
}

package practica

import (
	"fmt"
	"math"
	"strconv"
)

func fizzBuzz(num int) string {
	var s string = ""
	if math.Mod(float64(num), float64(3)) == 0 {
		s = "Fizz"
	}
	if math.Mod(float64(num), float64(5)) == 0 {
		s += "Buzz"
	}
	return s
}

func Ej6() {
	// Ejercicio FizzBuzz en su forma clásica es el siguiente: "Para cada número del 1 al 100, imprime 'Fizz' si el número es divisible por 3,
	// 'Buzz' si es divisible por 5 y 'FizzBuzz' si es divisible por ambos. Si no es divisible ni por 3 ni por 5, simplemente imprime el número."
	fmt.Print("EJERCICIO 6\n\n")

	fmt.Print("Impriendo del 1 al 100 con FizzBuzz\n\n")
	for i := 1; i <= 100; i++ {
		res := fizzBuzz((i))
		if (res == ""){
			res = strconv.Itoa(i);
		}
		fmt.Println(res)
	}
}

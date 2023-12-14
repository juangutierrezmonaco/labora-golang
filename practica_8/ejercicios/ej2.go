package practica

import (
	"fmt"
	"sync"
	"time"
)

func Ej2() {
	// Usando la técnica del sleep sort, hagamos un programa que imprima los primeros 10 números enteros en orden ascendiente.
	// ¿ Es posible hacerlo en orden descendiente?
	fmt.Print("EJERCICIO 2\n\n")

	n := 10
	upwards := false // Asc: true. Desc: false
	wg := sync.WaitGroup{}
	wg.Add(n)

	if upwards {
		fmt.Print("Los primeros 10 números enteros en orden ascendiente son: ")
	} else {
		fmt.Print("Los primeros 10 números enteros en orden descendiente son: ")
	}

	for i := 1; i <= n; i++ {
		go func(i int) {
			defer wg.Done()

			sleepingTime := i
			if !upwards {
				sleepingTime = n - i
			}
			time.Sleep(time.Duration(sleepingTime) * time.Millisecond)
			fmt.Printf("%d ", i)

		}(i)
	}
	fmt.Println()
	wg.Wait()

}

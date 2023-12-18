package practica

import (
	"fmt"
	"sync"
)

type SynchronizedMap[K comparable, V any] struct {
	mutex sync.Mutex
	m     map[K]V
}

func (sm *SynchronizedMap[K, V]) Set(key K, value V) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	sm.m[key] = value
}

func (sm *SynchronizedMap[K, V]) Get(key K) V {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	return sm.m[key]
}

func (sm *SynchronizedMap[K, V]) ToString() string {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	var str string
	fmt.Println("   Map   ")
	fmt.Println("----------")
	for key, value := range sm.m {
		str += fmt.Sprintf("-  \"%v\": %v\n", key, value)
	}
	return str
}

func NewSynchronizedMap[K comparable, V any]() SynchronizedMap[K, V] {
	return SynchronizedMap[K, V]{m: make(map[K]V)}
}

func Ej1() {
	// Entiendo el problema que se da cuando se quiere acceder de forma concurrente para modificar un mapa (tal como lo describe este ejemplo
	// (https://go.dev/play/p/ZT6Yb_6KGes)). Se desea implementar un tipo de datos abstracto (puede ser una struct) que represente un mapa sincronizado,
	// es decir, que tenga accesos concurrentes seguros o que permite operaciones concurrentes en forma segura (thread safe).
	// 		1. ¿Es importante que sea concurrentemente seguro a nivel solo lectura?
	// 		2. ¿Qué diferencia observas si usas métodos mutadores que sean pointer receiver y value receiver? Los métodos mutadores
	// 		son los que realizan cambios en el estado.
	fmt.Print("EJERCICIO 1\n\n")

	var wg sync.WaitGroup
	const n = 10000
	wg.Add(n)

	fmt.Printf("DEPOSITANDO %d VECES 1 PESO\n\n", n)

	moneyByName := NewSynchronizedMap[string, int]()
	name := "pepe"
	// Simulando el problema de los cajeros, le depositamos n veces 1 peso
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			currentValue := moneyByName.Get(name)
			moneyByName.Set(name, currentValue+1)
		}()
	}

	wg.Wait()
	fmt.Println("Protegiendo el mapa únicamente, el mapa queda: ")
	fmt.Println(moneyByName.ToString())

	// Hasta acá vemos que no hay ningún problema con la escritura, es safe, pero seguimos teniendo el problema de que
	// como la gorutinas se ejecutan en el orden que las agarre el planificador, sí o sí hay que trabar la operación de "depositar"
	mutex := sync.Mutex{}
	wg.Add(n)

	moneyByName.Set(name, 0)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			currentValue := moneyByName.Get(name)
			moneyByName.Set(name, currentValue+1)
			mutex.Unlock()
		}()
	}

	wg.Wait()	
	fmt.Println("Protegiendo el mapa y la operación de \"depósito\", el mapa queda: ")
	fmt.Println(moneyByName.ToString())

}

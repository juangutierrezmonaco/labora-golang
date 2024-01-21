package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labora/labora-golang/practica_13_servers/ej1/numbers"
	"github.com/labora/labora-golang/util"
)

type NumBody struct {
	Num int `json:"num"`
}

func handlerSum(w http.ResponseWriter, r *http.Request, isEvenSum bool) {
	var sumFunc func(int) int
	var typeText string
	if isEvenSum {
		sumFunc = numbers.SumNEvenNumbersGauss
		typeText = "even"
	} else {
		sumFunc = numbers.SumNOddNumbersGauss
		typeText = "odd"
	}

	if r.Method == "GET" {
		numStr := r.URL.Query().Get("number")
		num, err := util.Atoi(numStr)
		if err != nil {
			http.Error(w, "Invalid param, expecting a natural number.", 500)
			return
		}

		sum := sumFunc(num)

		w.Write([]byte(fmt.Sprintf("The sum of the first %d %s numbers is equal to:  %d.", num, typeText, sum)))
	}
	if r.Method == "POST" {
		var body NumBody
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, "Invalid param, expecting a natural number.", 500)
			return
		}

		num := body.Num
		sum := sumFunc(num)

		w.Write([]byte(fmt.Sprintf("The sum of the first %d %s numbers is equal to:  %d.", num, typeText, sum)))
	}
}

func handlerEvens(w http.ResponseWriter, r *http.Request) {
	handlerSum(w, r, true)
}

func handlerOdds(w http.ResponseWriter, r *http.Request) {
	handlerSum(w, r, false)
}

func main() {
	// Asumiendo que realizó los ejercicios para sumar los primeros n números pares y los primeros n números impares. Les voy a pedir que hagan lo siguiente:

	// 1. Crear un proyecto golang donde las funciones sean parte de un paquete que no sea el main (Por ejemplo: puede llamarlo numbers) y
	// que tenga un module path que sea un URL apuntando a un repositorio accesible.
	// 2. Crearse otro proyecto golang para escribir un servidor que tenga dos endpoints
	// 		a. /api/v1/sumEvens: Recibe como parámetro un número n y suma los primeros “n” números pares y devuelve en el cuerpo de la respuesta el resultado
	// 		de dicha suma.
	// 		b. /api/v1/sumOdds: Recibe como parámetro un número n y suma los primeros “n” números impares y devuelve en el cuerpo de la respuesta el resultado
	//		de dicha suma.

	// Se pide intencionalmente que NO se escriban las rutinas de código dentro del proyecto del servidor, sino que se importen usando los
	// mecanismos de dependencias que tiene el lenguaje.

	// 3. Se pide hacer ambos endpoints para recibir peticiones GET y POST, entender que los parámetros de la petición “viajan” de distinta forma si
	// es un GET o si es un POST. Si es un GET son parte de la URL de la petición y si es un POST es parte del cuerpo de la petición.

	// Si te parece más sencillo podés tener 4 endpoints, 2 para GET y 2 para POST.
	fmt.Print("EJERCICIO 1\n\n")

	PORT := ":8080"
	log.Print("Running server on " + PORT)
	http.HandleFunc("/api/v1/sumEvens", handlerEvens)
	http.HandleFunc("/api/v1/sumOdds", handlerOdds)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

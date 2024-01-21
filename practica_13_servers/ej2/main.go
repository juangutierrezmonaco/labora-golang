package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Num int

var num Num

func (num *Num) increment() {
	*num++
}

func (num *Num) decrement() {
	*num--
}

func (num *Num) get() int {
	return int(*num)
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inc", r.RemoteAddr)
	num.increment()
	w.WriteHeader(http.StatusOK)
}

func decrementCounter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Dec", r.RemoteAddr)
	num.decrement()
	w.WriteHeader(http.StatusOK)
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get", r.RemoteAddr)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(num.get())))
}

func main() {
	// Escribir un programa que levante un servidor en el puerto 9090. El programa debe mantener en memoria una variable entera definida a nivel “global”. Además configurar el servidor para tener dos endpoints.

	// 1. POST /api/v1/inc: Incrementa la variable entera en uno
	// 2. POST /api/v1/dec: Decrementa la variable entera en uno
	// 3. GET /api/v1/curr: Devuelve el valor de la variable entera.

	// Una vez escrito este programa servidor, vas a escribir un programa que juegue como cliente y tenga dos gorutinas. En una vas a enviar 10000 peticiones para incrementar y otra 1000 peticiones para decrementar. Finalmente vas a mandar una petición para obtener el valor de la variable entera una vez que terminen las dos gorutinas.

	// Se pide ejecutar el programa servidor y luego :

	// 1. Ejecutando una vez el programa cliente
	// 2. Ejecuta varias instancias del mismo programa cliente, es decir, ejecútalo varias veces en simultáneo.

	// Por último, se podrá apreciar que en el segundo caso para algún cliente el resultado que se le devuelve no es cero. ¿Esto a qué se debe? ¿ que hay que modificar para que a ambos clientes la petición “final” de obtener el valor de la variable entera devuelva un cero?

	// **Resoluciones posibles** *(if ❌🧠 then !👀, prohibido mirar sin antes pensar)*:

	// [int_server.go](https://go.dev/play/p/4AMPOqaM7pr)

	// [int_cliente.go](https://go.dev/play/p/DZR_R88SprX)
	fmt.Print("EJERCICIO 2\n\n")

	PORT := ":9090"
	log.Print("Running server on " + PORT)
	http.HandleFunc("/api/v1/inc", incrementCounter)
	http.HandleFunc("/api/v1/dec", decrementCounter)
	http.HandleFunc("/api/v1/curr", getCounter)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

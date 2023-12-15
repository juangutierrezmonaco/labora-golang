package practica

import "fmt"

func transformToSeconds(days int, hours int, minutes int, seconds int) int {
	secondsInMinute := 60
	secondsInHour := 60 * secondsInMinute
	secondsInDay := 24 * secondsInHour

	return seconds + minutes*secondsInMinute + hours*secondsInHour + days*secondsInDay
}

func Ej4() {
	// Redactar un algoritmo para pasar un periodo expresado en días, horas, minutos y segundos a periodo expresado en segundos.
	fmt.Print("EJERCICIO 4\n\n")

	var days, hours, minutes, seconds int
	fmt.Print("Ingrese los días: ")
	fmt.Scan(&days)
	fmt.Print("Ingrese las horas: ")
	fmt.Scan(&hours)
	fmt.Print("Ingrese los minutos: ")
	fmt.Scan(&minutes)
	fmt.Print("Ingrese los segundos: ")
	fmt.Scan(&seconds)

	resultInSeconds := transformToSeconds(days, hours, minutes, seconds)
	fmt.Printf("%d días, %d horas, %d minutos y %d segundos son: %d segundos.\n", days, hours, minutes, seconds, resultInSeconds)
}

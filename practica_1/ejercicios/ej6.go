package practica

import (
	"fmt"
	"math"
)

func transformToDaysHoursMinutesSeconds(secondsEntered float64) (float64, float64, float64, float64) {
	secondsInMinute := 60
	secondsInHour := 60 * secondsInMinute
	secondsInDay := 24 * secondsInHour

	var days, hours, minutes, seconds float64

	days, hours = math.Modf(secondsEntered / float64(secondsInDay))
	hours, minutes = math.Modf(hours * 24)
	minutes, seconds = math.Modf(minutes * 60)
	seconds = seconds * 60

	return days, hours, minutes, seconds
}

func Ej6() {
	// 6. Para valientes: Pasar un periodo expresado en segundos a un periodo expresado en días, horas, minutos y segundos.
	// a. 1030 segundos son 17 minutos con 10 segundos
	// b. 12045 segundos son 3 horas, 20 minutos con 45 segundos
	// c. 176520 segundos son 2 días, 1 hora, 2 minutos con 0 segundos.

	fmt.Print("EJERCICIO 6\n\n")

	var secondsEntered float64
	fmt.Print("Ingrese los segundos: ")
	fmt.Scan(&secondsEntered)

	days, hours, minutes, seconds := transformToDaysHoursMinutesSeconds(secondsEntered)
	fmt.Printf("%.0f segundos son: %.0f días, %.0f horas, %.0f minutos y %.0f segundos.\n", secondsEntered, days, hours, minutes, seconds)
}

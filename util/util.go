package util

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

/*** CONSOLE ***/
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearConsole() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

/*** MENU ***/
func validateAndSetOption(inf int, sup int) int {
	var op int
	fmt.Printf("Hola! Ingresá qué ejercicio querés correr (%d a %d, 0 para SALIR): ", inf+1, sup)
	fmt.Scan(&op)
	for op < inf || op > sup {
		fmt.Printf("Opción incorrecta: Ingrese nuavemente (%d a %d, 0 para SALIR): ", inf+1, sup)
		fmt.Scan(&op)
	}
	return op
}

func RunMenu(infOp int, supOp int, functions []func()) {
	op := validateAndSetOption(infOp, supOp)

	if op != 0 {

		functions[op-1]()

		var runAgain string
		fmt.Print("\n¿Querés correr volver al INICIO? S/s: ")
		fmt.Scan(&runAgain)
		if runAgain == "S" || runAgain == "s" {
			ClearConsole()
			RunMenu(infOp, supOp, functions)
		} else {
			fmt.Println("¡Adiós!")
		}
	}
}

/*** ARRAYS ***/
func JoinIntArr(arr []int, sep string) string {
	stringArr := make([]string, len(arr))
	for i, v := range arr {
		stringArr[i] = strconv.Itoa(v)
	}

	return strings.Join(stringArr, sep)
}

/*** I/O ***/
func ScanSentence() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

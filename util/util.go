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

/*** MENU FOR EXERCISES ***/
func validateAndSetOption(inf int, sup int) int {
	var op int
	fmt.Printf("Ingresá qué ejercicio querés correr (%d a %d, 0 para SALIR): ", inf+1, sup)
	fmt.Scan(&op)
	for op < inf || op > sup {
		fmt.Printf("Opción incorrecta: Ingrese nuavemente (%d a %d, 0 para SALIR): ", inf+1, sup)
		fmt.Scan(&op)
	}
	return op
}

func RunMenu(infOp int, supOp int, functions []func()) {
	var op int
	if len(functions) > 1 {
		op = validateAndSetOption(infOp, supOp)
	} else {
		op = 1
	}

	if op != 0 {

		functions[op-1]()

		var runAgain string
		fmt.Print("\n¿Querés volver al INICIO? S/s: ")
		fmt.Scan(&runAgain)
		if runAgain == "S" || runAgain == "s" {
			ClearConsole()
			RunMenu(infOp, supOp, functions)
		} else {
			fmt.Println("¡Adiós!")
		}
	}
}

/*** GENERIC MENU ***/
type Menu struct {
	functions []func()
	options   []string
	title     string
}

func NewMenu(functions []func(), options []string, title string) Menu {
	var m Menu
	m.functions = functions
	m.options = options
	m.options = append(m.options, "Volver")
	m.title = title
	return m
}

func (m *Menu) printTitle() {
	fmt.Printf("======= %s =======\n\n", m.title)
}

func (m *Menu) Print() {
	m.printTitle()
	for i, opt := range m.options {
		fmt.Printf("%d. %s\n", i+1, opt)
	}
	fmt.Println()
}

func (m *Menu) validateAndSetOption() int {
	var op int
	inf := 1
	sup := len(m.options)

	fmt.Print("Ingresá una opción: ")
	fmt.Scan(&op)
	for op < inf || op > sup {
		ClearConsole()
		m.Print()
		fmt.Printf("Opción incorrecta: Ingrese nuavemente (%d a %d): ", inf, sup)
		fmt.Scan(&op)
	}
	return op
}

func (m *Menu) Run() int {
	ClearConsole()

	// Printing menu
	m.Print()

	// Choose an option
	op := m.validateAndSetOption()
	ClearConsole()

	if op != len(m.functions)+1 && m.functions != nil {

		m.functions[op-1]()

		var runAgain string
		fmt.Print("\n¿Querés volver al INICIO? S/s: ")
		fmt.Scan(&runAgain)
		if runAgain == "S" || runAgain == "s" {
			ClearConsole()
			m.Run()
		} else {
			fmt.Println("¡Adiós!")
		}
	}
	return op
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

/*** CASTING ***/
func Itoa(num int) string {
	return strconv.Itoa(num)
}

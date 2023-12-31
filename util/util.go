package util

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
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
		WaitForEnter()
		ClearConsole()
		m.Run()
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

/*** I/O & Formatting ***/
func ConfirmAction(text string) bool {
	fmt.Printf("%s S/s: ", text)
	var char string
	fmt.Scan(&char)

	doAction := char == "S" || char == "s"
	return doAction
}

func TryAgainWhenError(err error, funcToRunAgain func()) {
	fmt.Printf("\nHubo un error: %s\n", err.Error())
	tryAgain := ConfirmAction("¿Quiere intentar nuevamente?")
	if tryAgain {
		funcToRunAgain()
	}
}

func ScanSentence() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func ScanValue(value interface{}) (interface{}, error) {
	input := ScanSentence()

	switch value.(type) {
	case *bool:
		boolValue, err := strconv.ParseBool(strings.TrimSpace(input))
		if err != nil {
			return nil, err
		}
		return boolValue, nil
	case *float64:
		floatValue, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			return nil, err
		}
		return floatValue, nil
	case *float32:
		floatValue, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
		if err != nil {
			return nil, err
		}
		return float32(floatValue), nil
	case *int:
		intValue, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			return nil, err
		}
		return intValue, nil
	default:
		return nil, fmt.Errorf("Tipo no admitido")
	}
}

func WaitForEnter() {
	fmt.Print("Presioná 'Enter' para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func TextWithFiller(filler string, quantityPerSide int, target string) string {
	fillerText := strings.Repeat(filler, quantityPerSide)
	return fmt.Sprintf("%s %s %s", fillerText, target, fillerText)
}

func CapitalizeWords(str string) string {
	return strings.Title(strings.ToLower(str))
}

/*** CASTING ***/
func Itoa(num int) string {
	return strconv.Itoa(num)
}

func Atoi(str string) (int, error) {
	return strconv.Atoi(str)
}

func StrToFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

/*** GENERATING STUFF ***/
func GenerateRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

package util

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func GenerateRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func WaitForEnter() {
	fmt.Print("Presioná 'Enter' para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func DeleteLinesFromTerminal(n int) {
	fmt.Printf("\033[%dA\033[J", n)
}

func PressEnterToContinue() {
	WaitForEnter()
	DeleteLinesFromTerminal(1)
}

func TextWithFiller(filler string, quantityPerSide int, target string) string {
	fillerText := strings.Repeat(filler, quantityPerSide)
	return fmt.Sprintf("%s %s %s", fillerText, target, fillerText)
}

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

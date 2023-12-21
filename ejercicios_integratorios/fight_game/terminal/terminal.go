package terminal

import (
	"fmt"
	"time"
	"unicode"

	"github.com/labora/labora-golang/ejercicios_integratorios/fight_game/contender"
	"github.com/labora/labora-golang/ejercicios_integratorios/fight_game/util"
)

func PrintTimer(seconds int) {
	fmt.Print("\nSiguiente ronda en: ")
	for i := 0; i < seconds; i++ {
		remainingSeconds := seconds - i
		fmt.Printf(" %d", remainingSeconds)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n\n")
}

func PrintTextSlowly(text string, milliseconds int) {
	var wordCount int
	for _, letter := range text {
		fmt.Printf("%s", string(letter))
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)

		// Prints an enter every 70 lines when it founds an space
		if unicode.IsSpace(letter) {
			wordCount++
		}
		// Using \t as \n to simplify the logic
		if wordCount == 10 || string(letter) == "\t" {
			fmt.Println()
			wordCount = 0
		}
	}
	fmt.Printf("\n\n")
}

func PrintIntro(contenders []contender.Contender) {
	intro := "En NeoEclipsis, una ciudad sumida en la oscuridad y la corrupción, tres figuras emergen para enfrentarse en una lucha épica. Sir Luminis, el Paladín dedicado a la luz; el Teniente Vanguard, un policía implacable; y Sombra Nocturna, un criminal astuto.\n\tLa ciudad, una vez utópica, se tambalea al borde del colapso debido al crimen desenfrenado. Estos tres personajes, cada uno con su visión única de justicia, chocan en \"Sombras en Guerra\". ¿Prevalecerá la luz sobre la oscuridad o las sombras engullirán la ciudad? \t¡Prepárate para una lucha que cambiará el destino de NeoEclipsis para siempre!"
	PrintTextSlowly(intro, 15)

	fmt.Printf("¡Estos son nuestros valientes contendientes!, ¿quién ganará? ¡Sólo la SUERTE podrá decidirlo!\n\n")
	for _, c := range contenders {
		fmt.Printf("- %s\n", c.ToString())
	}
	fmt.Println()
}

func PrintPreCombatInfo(round int, attackerName string, dialogue string) {
	fmt.Printf("\n\nRONDA %d\n\n", round)
	fmt.Printf("* Se lanzó un dado de tres caras y el destino ha determinado complacer al %s\n", attackerName)
	time.Sleep(1 * time.Second)
	fmt.Printf("\n\t* El %s dice: # %s #\n\n", attackerName, dialogue)
	time.Sleep(2 * time.Second)
}

func PrintResult(w contender.Contender) {
	PrintTextSlowly("\n\nEn la encrucijada del destino, un campeón se ha alzado sobre los demás...", 50)
	util.PressEnterToContinue()
	fmt.Printf("\n\nLA FORTUNA Y EL DESTINO HAN DETERMINADO QUE EL GANADOR ES EL %s\n\n", w.ToString())
}

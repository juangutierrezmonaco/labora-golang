package combat

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/labora/labora-golang/ejercicios_integratorios/fight_game/contender"
	"github.com/labora/labora-golang/ejercicios_integratorios/fight_game/terminal"
	"github.com/labora/labora-golang/ejercicios_integratorios/fight_game/util"
)

func rollATriangularDice() int {
	return util.GenerateRandomNumber(0, 2)
}

func rollDiceBetweenAliveContenders(contenders []contender.Contender) int {
	diceResult := rollATriangularDice()
	for !contenders[diceResult].IsAlive() {
		diceResult = rollATriangularDice()
	}

	return diceResult
}

func selectVictims(contenders []contender.Contender, attackerIndex int) []contender.Contender {
	var victims []contender.Contender
	for i, c := range contenders {
		if c.IsAlive() && i != attackerIndex {
			victims = append(victims, c)
		}
	}
	return victims
}

func chooseAttackIntensity(wg *sync.WaitGroup) int {
	wg.Add(2)
	var intensityFactor int
	go func(int) {
		defer wg.Done()
		defer wg.Done()
		intensityFactor = determineHitPower()
	}(intensityFactor)
	wg.Wait()
	return intensityFactor
}

func runCombat(attackIntensity int, attackerName string, victims []contender.Contender, attackInfo string) bool {
	survivingContenders := 0
	if attackIntensity == 0 {
		fmt.Printf("Parece que la suerte no favorece al %s, pues su ataque ha tenido intesidad %d.\n", attackerName, attackIntensity)
	} else {
		wg := sync.WaitGroup{}
		intensityFactor := chooseAttackIntensity(&wg)
		wg.Wait()

		fmt.Printf("El ataque de %d ha sido modificado por un factor de: ~%d%%\n", attackIntensity, intensityFactor)
		attackIntensity = int((float64(intensityFactor) / float64(100)) * float64(attackIntensity))
		fmt.Printf("Resultando en: %d\n\n", attackIntensity)

		fmt.Println(attackInfo)
		fmt.Printf("* El %s ataca a con intensidad * %d * a:\n", attackerName, attackIntensity)
		for _, v := range victims {
			fmt.Printf("- %s\n", v.ToString())
			v.RecieveAttack(attackIntensity)

			if v.IsAlive() {
				survivingContenders++
			}
		}
	}

	// Returns if someone from the victims is alive
	return attackIntensity != 0 && survivingContenders == 0
}

func waitForEnter(exitCh *chan bool) {
	var reader = bufio.NewReader(os.Stdin)

	select {
	case <-*exitCh:
		close(*exitCh)
	default:
		input, _, _ := reader.ReadRune()
		if input == 10 { // enter
			*exitCh <- true
		}
	}
}

func triggerCountdown(exitCh *chan bool, timerCh *chan time.Duration, seconds int) {
	defer close(*timerCh)
	finishTime := time.Duration(seconds) * time.Second
	timeLeft := finishTime

	millisecondsPassed := 0
	for timeLeft >= 0 {
		// wait for 0,1s
		time.Sleep(100 * time.Millisecond)

		// send the remaining seconds
		remainingSeconds := time.Duration(math.Ceil(float64(timeLeft.Seconds()))) * time.Second
		*timerCh <- time.Duration(remainingSeconds)
		timeLeft = finishTime - (time.Duration(100*millisecondsPassed) * time.Millisecond)

		// Increase how many milliseconds passed
		millisecondsPassed++
	}

	// Send a value to exit
	*exitCh <- true
}

func determineHitPower() int {
	exitCh := make(chan bool)
	timerCh := make(chan time.Duration)
	go waitForEnter(&exitCh)
	go triggerCountdown(&exitCh, &timerCh, 15)

	barLength := 30
	hitIntensity := 1
	firstTime := true
	for {
		select {
		case <-exitCh:
			porcentualHitIntensity := (float64(hitIntensity) / float64(barLength)) * 100

			// Remove the line with the 'Press enter to continue'
			util.DeleteLinesFromTerminal(3)
			fmt.Println()
			return int(porcentualHitIntensity)
		case timeLeft := <-timerCh:
			if !firstTime {
				util.DeleteLinesFromTerminal(4)
			} else {
				firstTime = false
			}
			str := fmt.Sprintf("[ %s🥊 ", strings.Repeat("=", hitIntensity))
			fmt.Printf("%s %*s\n", str, barLength-hitIntensity, "]")
			fmt.Printf("\n\nPresioná 'ENTER' para determinar la fuerza del golpe...(¡Te quedan %v segundos!)\n", timeLeft)
			hitIntensity++
			hitIntensity %= barLength + 1
		}
	}
}

func RunGame() {
	util.ClearConsole()
	contenders := []contender.Contender{
		contender.NewPolice(),
		contender.NewCriminal(),
		contender.NewPaladin(),
	}

	// terminal.PrintIntro(contenders)
	util.PressEnterToContinue()

	noVictimSurvived := false
	rounds := 1
	var winner contender.Contender

	for !noVictimSurvived {
		// Roll the dice
		diceResult := rollDiceBetweenAliveContenders(contenders)

		// Select the attacker and the victims
		attacker := contenders[diceResult]
		attackerName := contender.FormatContenderName(attacker)
		victims := selectVictims(contenders, diceResult)

		// Display who's gonna attack who
		terminal.PrintPreCombatInfo(rounds, attackerName, attacker.GetAttackDialogue())

		// Get's attack intensity
		attackIntensity, attackInfo := attacker.ThrowAttack()
		noVictimSurvived = runCombat(attackIntensity, attackerName, victims, attackInfo)

		// Display current results
		fmt.Printf("\nEstos son los resultados hasta el momento:\n")
		for _, c := range contenders {
			fmt.Printf("- %s\n", c.ToString())
		}

		// Next round
		if noVictimSurvived {
			winner = attacker
		} else {
			// Wait 3 seconds for next combat
			terminal.PrintTimer(5)
			rounds++
		}
	}

	terminal.PrintResult(winner)
}

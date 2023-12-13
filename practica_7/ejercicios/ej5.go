package practica

import (
	"fmt"
	"math"
	_ "math/rand"
	"strings"
	"time"
	_ "time"
	"unicode"

	"github.com/labora/labora-golang/util"
)

/***************************************************************/
/*                            HUMAN                            */
/***************************************************************/
type Human struct {
	Life int //0..100
	Role string
}

func (h *Human) GetLife() int {
	return h.Life
}

func (h *Human) SetLife(life int) {
	if life < 0 {
		h.Life = 0
	} else {
		h.Life = life
	}
}

func (h *Human) GetRole() string {
	return h.Role
}

func (h *Human) SetRole(role string) {
	h.Role = role
}

func (h *Human) IsAlive() bool {
	return h.Life > 0
}

func NewHuman(life int, role string) *Human {
	human := Human{Life: life, Role: role}
	return &human
}

/***************************************************************/
/*                          CONTENDER                          */
/***************************************************************/
type Contender interface {
	ThrowAttack() int
	RecieveAttack(intensity int)
	GetAttackDialogue() string
	IsAlive() bool
	GetLife() int
	GetRole() string
	ToString() string
}

const (
	policeInitialLife   = 100
	criminalInitialLife = 100
	paladinInitialLife  = 200
)

/***************************************************************/
/*                           POLICE                            */
/***************************************************************/
type Police struct {
	Human
	Armour int // 0..50
}

func NewPolice() *Police {
	return &Police{
		Human:  *NewHuman(policeInitialLife, "Policía"),
		Armour: 50,
	}
}

func (_ *Police) ThrowAttack() int {
	return util.GenerateRandomNumber(0, 4)
}

func (p *Police) RecieveAttack(intensity int) {
	if p.Armour > 0 {
		diff := (p.Armour - intensity)
		armorCoversWholeAttack := diff > 0
		if armorCoversWholeAttack {
			p.Armour -= intensity
			intensity = 0
		} else {
			intensity -= p.Armour
			p.Armour = 0
		}
	}
	p.SetLife(p.Life - intensity)
}

func (_ *Police) GetAttackDialogue() string {
	lines := []string{
		"PREPÁRATE PARA TU ARRESTO.",
		"EN LA LEY ENCUENTRO MI FUERZA, Y HOY, TE HACERÉ SENTIR SU PESO.",
		"EN NOMBRE DE LA JUSTICIA, TE DETENDRÉ.",
		"CUMPLIRÉ MI DEBER, CUESTE LO QUE CUESTE.",
		"LA JUSTICIA NO TIENE PIEDAD, Y YO SOY SU EJECUTOR.",
	}
	selectedLine := util.GenerateRandomNumber(0, 4)
	return lines[selectedLine]
}

func (p *Police) ToString() string {
	var str string
	if p.IsAlive() {
		str = fmt.Sprintf("%s (Vida = %d, Armadura = %d)", formatContenderName(p), p.Life, p.Armour)
	} else {
		str = fmt.Sprintf("El %s ya no se encuentra en el reino de los vivos.", p.Role)
	}
	return str
}

/***************************************************************/
/*                          CRIMINAL                           */
/***************************************************************/
type Criminal struct {
	Human
}

func NewCriminal() *Criminal {
	return &Criminal{
		Human: *NewHuman(criminalInitialLife, "Criminal"),
	}
}

func (_ *Criminal) ThrowAttack() int {
	return util.GenerateRandomNumber(0, 8)
}

func (c *Criminal) RecieveAttack(intensity int) {
	c.SetLife(c.Life - intensity)
}

func (_ *Criminal) GetAttackDialogue() string {
	lines := []string{
		"LAS CADENAS DE LA MORALIDAD NO ME DETENDRÁN.",
		"CADA CRIMEN ES UNA OBRA MAESTRA, Y ESTA SERÁ TU FINAL.",
		"EL CRIMEN PAGA, PERO HOY PAGARÁS TÚ.",
		"MI MALDAD ES INCONTENIBLE, Y HOY SERÁS TESTIGO DE SU FURIA.",
		"EL MIEDO ES MI HERRAMIENTA, Y HOY LO FORJARÉ EN TU MENTE.",
	}
	selectedLine := util.GenerateRandomNumber(0, 4)
	return lines[selectedLine]
}

func (c *Criminal) ToString() string {
	var str string
	if c.IsAlive() {
		str = fmt.Sprintf("%s (Vida = %d)", formatContenderName(c), c.Life)
	} else {
		str = fmt.Sprintf("El %s ya no se encuentra en el reino de los vivos.", c.Role)
	}
	return str
}

/***************************************************************/
/*                          PALADIN                            */
/***************************************************************/
type Paladin struct {
	Human
}

func NewPaladin() *Paladin {
	return &Paladin{
		Human: *NewHuman(paladinInitialLife, "Paladín"),
	}
}

func (p *Paladin) ThrowAttack() int {
	damageFactor := float64(p.Life) / float64(paladinInitialLife)
	attackIntensity := util.GenerateRandomNumber(0, 10)
	resultingDamage := damageFactor * float64(attackIntensity)
	return int(math.Round(resultingDamage))
}

func (p *Paladin) RecieveAttack(intensity int) {
	p.SetLife(p.Life - intensity)
}

func (_ *Paladin) GetAttackDialogue() string {
	lines := []string{
		"CON LA LUZ COMO MI GUÍA, TE DERROTARÉ.",
		"EN NOMBRE DE LA VIRTUD, CAERÁS.",
		"PREPÁRATE PARA LA PURIFICACIÓN, PUES MI IRA NO CONOCE LÍMITES.",
		"EL MAL NO ENCONTRARÁ REFUGIO ANTE MI LUZ.",
		"POR LA GLORIA Y EL HONOR, TE VENCERÉ.",
	}
	selectedLine := util.GenerateRandomNumber(0, 4)
	return lines[selectedLine]
}

func (p *Paladin) ToString() string {
	var str string
	if p.IsAlive() {
		str = fmt.Sprintf("%s (Vida = %d)", formatContenderName(p), p.Life)
	} else {
		str = fmt.Sprintf("El %s ya no se encuentra en el reino de los vivos.", p.Role)
	}
	return str
}

// Helper methods
func printTimer(seconds int) {
	fmt.Print("\nSiguiente ronda en: ")
	for i := 0; i < seconds; i++ {
		remainingSeconds := seconds - i
		fmt.Printf(" %d", remainingSeconds)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n\n")
}

func waitForEnter() {
	if interactive {
		util.WaitForEnter()
	}
}

func formatContenderName(c Contender) string {
	return util.TextWithFiller("~", 1, strings.ToUpper(c.GetRole()))
}

func printTextSlowly(text string, milliseconds int) {
	var wordCount int
	for _, letter := range text {
		fmt.Printf("%s", string(letter))
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)

		// Prints an enter every 70 lines when it founds an space
		if unicode.IsSpace(letter) {
			wordCount++
		}
		if wordCount == 10 || string(letter) =="\t"{
			fmt.Println()
			wordCount=0
		}
	}
	fmt.Printf("\n\n")
}

func printIntro(contenders []Contender) {
	intro := "En NeoEclipsis, una ciudad sumida en la oscuridad y la corrupción, tres figuras emergen para enfrentarse en una lucha épica. Sir Luminis, el Paladín dedicado a la luz; el Teniente Vanguard, un policía implacable; y Sombra Nocturna, un criminal astuto.\n\tLa ciudad, una vez utópica, se tambalea al borde del colapso debido al crimen desenfrenado. Estos tres personajes, cada uno con su visión única de justicia, chocan en \"Sombras en Guerra\". ¿Prevalecerá la luz sobre la oscuridad o las sombras engullirán la ciudad? \t¡Prepárate para una lucha que cambiará el destino de NeoEclipsis para siempre!"
	printTextSlowly(intro, 15)

	fmt.Printf("¡Estos son nuestros valientes contendientes!, ¿quién ganará? ¡Sólo la SUERTE podrá decidirlo!\n\n")
	for _, c := range contenders {
		fmt.Printf("- %s\n", c.ToString())
	}
	fmt.Println()
}

func printPreCombatInfo(round int, attackerName string, dialogue string) {
	fmt.Printf("\n\nRONDA %d\n\n", round)
	fmt.Printf("* Se lanzó un dado de tres caras y el destino ha determinado complacer al %s\n", attackerName)
	waitForEnter()
	fmt.Printf("\t* El %s dice: # %s #\n\n", attackerName, dialogue)
}

func printResult(w Contender) {
	printTextSlowly("\n\nEn la encrucijada del destino, un campeón se ha alzado sobre los demás...", 50)
	waitForEnter()
	fmt.Printf("\n\nLA FORTUNA Y EL DESTINO HAN DETERMINADO QUE EL GANADOR ES EL %s\n\n", w.ToString())
}

func rollDiceBetweenAliveContenders(contenders []Contender) int {
	diceResult := rollATriangularDice()
	for !contenders[diceResult].IsAlive() {
		diceResult = rollATriangularDice()
	}

	return diceResult
}

func selectVictims(contenders []Contender, attackerIndex int) []Contender {
	var victims []Contender
	for i, c := range contenders {
		if c.IsAlive() && i != attackerIndex {
			victims = append(victims, c)
		}
	}
	return victims
}

func runCombat(attackIntensity int, attackerName string, victims []Contender) bool {
	survivingContenders := 0
	if attackIntensity == 0 {
		fmt.Printf("Parece que la suerte no favorece al %s, pues su ataque ha tenido intesidad %d.\n", attackerName, attackIntensity)
	} else {
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

// Deactivate the interactive features
const interactive = true

func Ej5() {
	// Ejercicio de contrincantes
	fmt.Print("EJERCICIO 5\n\n")
	util.ClearConsole()

	contenders := []Contender{
		NewPolice(),
		NewCriminal(),
		NewPaladin(),
	}

	printIntro(contenders)
	waitForEnter()

	noVictimSurvived := false
	rounds := 1
	var winner Contender

	for !noVictimSurvived {
		// Roll the dice
		diceResult := rollDiceBetweenAliveContenders(contenders)

		// Select the attacker and the victims
		attacker := contenders[diceResult]
		attackerName := formatContenderName(attacker)
		victims := selectVictims(contenders, diceResult)

		// Display who's gonna attack who
		printPreCombatInfo(rounds, attackerName, attacker.GetAttackDialogue())

		// Get's attack intensity
		attackIntensity := attacker.ThrowAttack()
		noVictimSurvived = runCombat(attackIntensity, attackerName, victims)

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
			printTimer(3)
			rounds++
		}
	}

	printResult(winner)
}

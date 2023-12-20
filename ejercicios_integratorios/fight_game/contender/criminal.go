package contender

import (
	"fmt"

	"github.com/labora/labora-golang/ejercicios_integratorios/fight_game/util"
)

type Criminal struct {
	Human
}

func NewCriminal() *Criminal {
	return &Criminal{
		Human: *NewHuman(criminalInitialLife, "Criminal"),
	}
}

func (_ *Criminal) ThrowAttack() (int, string) {
	return util.GenerateRandomNumber(0, 80), ""
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
		str = fmt.Sprintf("%s (Vida = %d)", FormatContenderName(c), c.Life)
	} else {
		str = fmt.Sprintf("El %s ya no se encuentra en el reino de los vivos.", c.Role)
	}
	return str
}

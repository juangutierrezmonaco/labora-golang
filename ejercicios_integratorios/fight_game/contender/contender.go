package contender

import (
	"strings"

	"github.com/labora/labora-golang/ejercicios_integratorios/fight_game/util"
)

type Contender interface {
	ThrowAttack() (int, string) // The second argument is the explanation if something affects it's attack
	RecieveAttack(intensity int)
	GetAttackDialogue() string
	IsAlive() bool
	GetLife() int
	GetRole() string
	ToString() string
}

func FormatContenderName(c Contender) string {
	return util.TextWithFiller("~", 1, strings.ToUpper(c.GetRole()))
}

const (
	policeInitialLife   = 1000
	criminalInitialLife = 1000
	paladinInitialLife  = 2000
)

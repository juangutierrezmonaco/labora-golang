package contender

import (
	"fmt"
	"math"

	"github.com/labora/labora-golang/ejercicios_integratorios/fight_game/util"
)

type Paladin struct {
	Human
}

func NewPaladin() *Paladin {
	return &Paladin{
		Human: *NewHuman(paladinInitialLife, "Paladín"),
	}
}

func (p *Paladin) ThrowAttack() (int, string) {
	damageFactor := float64(p.Life) / float64(paladinInitialLife)
	attackIntensity := util.GenerateRandomNumber(0, 100)
	resultingDamage := damageFactor * float64(attackIntensity)
	damageFactorPorcentual := damageFactor * 100
	str := ""
	if damageFactor != 1 {
		str = fmt.Sprintf("Al tener un ~%.0f%% (%d/%d) de vida, el ataque del %s sólo afecta en un ~%.0f%% (Ataque pasa de %d a %d):", damageFactorPorcentual, p.Life, paladinInitialLife, p.GetRole(), damageFactorPorcentual, attackIntensity, int(math.Round(resultingDamage)))
	}
	return int(math.Round(resultingDamage)), str
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
		str = fmt.Sprintf("%s (Vida = %d)", FormatContenderName(p), p.Life)
	} else {
		str = fmt.Sprintf("El %s ya no se encuentra en el reino de los vivos.", p.Role)
	}
	return str
}

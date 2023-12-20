package contender

import (
	"fmt"

	"github.com/labora/labora-golang/ejercicios_integratorios/fight_game/util"
)

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

func (_ *Police) ThrowAttack() (int, string) {
	return util.GenerateRandomNumber(0, 40), ""
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
		str = fmt.Sprintf("%s (Vida = %d, Armadura = %d)", FormatContenderName(p), p.Life, p.Armour)
	} else {
		str = fmt.Sprintf("El %s ya no se encuentra en el reino de los vivos.", p.Role)
	}
	return str
}

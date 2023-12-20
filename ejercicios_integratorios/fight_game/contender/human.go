package contender

type Human struct {
	Life int //0..1000
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

package main

type Player struct {
}

func (p *Player) Attack(t *BeAttackable) {
	// ...
}

type Monster struct {
}

func (m *Monster) Attack(t *BeAttackable) {
	// ...
	t.BeAttacked()
}

type Attackable interface {
	Attack(BeAttackable)
}

type BeAttackable interface {
	BeAttacked()
}

func Attack(attacker *Attackable, defender *BeAttackable) {
	attacker.Attack(defender)
}

// if add chest ...

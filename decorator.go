package main

import (
	"fmt"
)

type IWeapon interface {
	getAtk() int
	getDur() int
	getMp() int
}

// Concrete weapon
type Sword struct{}

func (s Sword) getAtk() int {
	return 10
}

func (s Sword) getDur() int {
	return 5
}

func (s Sword) getMp() int {
	return 0
}

// Decorator for Attack
type Sharper struct {
	weapon IWeapon
}

func (s Sharper) getAtk() int {
	return s.weapon.getAtk() + 5
}

func (s Sharper) getDur() int {
	return s.weapon.getDur()
}

func (s Sharper) getMp() int {
	return s.weapon.getMp()
}

// Decorator for Durability
type Harder struct {
	weapon IWeapon
}

func (h Harder) getAtk() int {
	return h.weapon.getAtk()
}

func (h Harder) getDur() int {
	return h.weapon.getDur() + 10
}

func (h Harder) getMp() int {
	return h.weapon.getMp()
}

// Decorator for Magic power
type Enchanter struct {
	weapon IWeapon
}

func (e *Enchanter) getAtk() int {
	return e.weapon.getAtk()
}

func (e *Enchanter) getDur() int {
	return e.weapon.getDur()
}

func (e *Enchanter) getMp() int {
	return e.weapon.getMp() + 5
}

func showStats(w IWeapon) {
	fmt.Printf("Attack: %d \nDurability: %d \nMagic power: %d \n\n", w.getAtk(), w.getDur(), w.getMp())
}

func main() {
	Sword := Sword{}
	showStats(Sword)

	SharperSword := Sharper{Sword}
	showStats(SharperSword)

	MoreSharperSword := Sharper{SharperSword}
	showStats(MoreSharperSword)

	MoreSharperAndDurableSword := Harder{MoreSharperSword}
	showStats(MoreSharperAndDurableSword)

	MoreSharperAndDurableMagicSword := Enchanter{MoreSharperAndDurableSword}
	showStats(&MoreSharperAndDurableMagicSword)
}

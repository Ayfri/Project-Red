package main

import (
	"fmt"
)

type Character struct {
	name      string
	class     string
	lvl       int
	maxHealth int
	health    int
	skill     []string
	money     int
	equipment Equipment
	inventory Inventory
}

type Equipment struct {
	head  Item
	tunic Item
	boots Item
}

type EquipmentType int

const (
	Head EquipmentType = iota
	Tunic
	Boots
)

func (character Character) dead() {
	if character.health >= 0 {
		fmt.Printf("You're dead.")
		character.health = character.maxHealth / 2
		fmt.Printf("Resurrected with %d", character.health)
	}
}

func (character *Character) showHealth() {
	fmt.Printf("Health : %v/%v\n", character.health, character.maxHealth)
}

func (character Character) spellBook(name string) {
	if Contains(make([]interface{}, len(character.skill)), name) {
		fmt.Printf("You already have the spell '%v'.", name)
		return
	}
	character.skill = append(character.skill, "Fireball")
	character.inventory.removeItem(name, 1)
}

func (character Character) displayInfo() {
	fmt.Printf(
		`
Name: %v
Class: %v
Health: %v/%v
Lvl: %v
Money: %v
`,
		character.name,
		character.class,
		character.health,
		character.maxHealth,
		character.lvl,
		character.money,
	)
}

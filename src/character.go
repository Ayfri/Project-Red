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
	inventory Inventory
}

func (character Character) dead() {
	if character.health >= 0 {
		fmt.Printf("You're dead.")
		character.health = character.maxHealth / 2
		fmt.Printf("Resurrected with %d", character.health)
	}
}

func(character Character) spellBook(name string) {
	if Contains(make([]interface{}, len(character.skill)), name) {
		fmt.Printf("You already have the spell '%v'.", name)
		return
	}
	character.skill = append(character.skill, "Fireball")
	character.inventory.removeItem(name, 1)
}

func displayInfo(character Character) {
	fmt.Printf(
		`
Name: %v
Class: %v
Health: %v/%v
Lvl: %v
`,
		character.name,
		character.class,
		character.health,
		character.maxHealth,
		character.lvl,
	)
}

package main

import (
	"fmt"
)

func (character *Character) dead() {
	if character.health >= 0 {
		fmt.Printf("You're dead.")
		character.health = character.maxHealth / 2
		fmt.Printf("Resurrected with %d", character.health)
	}
}

func (character *Character) equip(item Item) {
	switch item.equipmentType {
	case Head:
		if character.equipment.head != nil {
			character.inventory.addItem(*character.equipment.head)
		}
		character.equipment.head = &item
	case Tunic:
		if character.equipment.tunic != nil {
			character.inventory.addItem(*character.equipment.tunic)
		}
		character.equipment.tunic = &item

	case Boots:
		if character.equipment.boots != nil {
			character.inventory.addItem(*character.equipment.boots)
		}
		character.equipment.boots = &item
	}
}

func (character *Character) showHealth() {
	fmt.Printf("Health : %v/%v\n", character.health, character.maxHealth)
}

func (character *Character) spellBook(name string) {
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
Equipment: %v
`,
		character.name,
		character.class,
		character.health,
		character.maxHealth,
		character.lvl,
		character.money,
		character.equipment.Show(),
	)
}

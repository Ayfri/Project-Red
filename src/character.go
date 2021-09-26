package main

import (
	"fmt"
	"github.com/fatih/color"
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
		colorFprintf("You already have the spell %v.", color.CyanString(name))
		return
	}
	character.skill = append(character.skill, "Fireball")
	character.inventory.removeItem(name, 1)
}

func (character *Character) displayInfo() {
	fmt.Fprintf(
		color.Output,
		`Name: %v
Class: %v
Health: %v
Lvl: %v
Money: %v
Equipment: %v
`,
		boldString(character.name),
		color.GreenString(character.class),
		color.RedString(fmt.Sprintf("%v/%v", character.health, character.maxHealth)),
		color.MagentaString(str(character.lvl)),
		color.YellowString(str(character.money)),
		character.equipment.Show(),
	)
}

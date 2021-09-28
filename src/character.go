package main

import (
	"fmt"
)

func (character *Character) attack(monster *Monster) {
	damages := 5
	monster.Health -= damages
	printAttack(*character, *monster, damages)
}

func (character *Character) dead() {
	if character.getHealth() >= 0 {
		fmt.Println("You're dead.")
		character.Health = character.MaxHealth / 2
		colorFprintf("Resurrected with %s health.\n", yellowString(character.showHealth()))
	}
}

func (character *Character) displayInfo() {
	colorFprintf(
		`Name: %v
Class: %v
Health: %v
Lvl: %v
Money: %v
Equipment: %v
`,
		boldString(character.Name),
		greenString(character.Class),
		redString(fmt.Sprintf("%v/%v", character.getHealth(), character.getMaxHealth())),
		magentaString(str(character.Lvl)),
		yellowString(str(character.Money)),
		character.Equipment.Show(),
	)
}

func (character *Character) equip(item Item) {
	switch item.EquipmentType {
	case Head:
		if character.Equipment.Head != nil {
			character.Inventory.addItem(*character.Equipment.Head)
		}
		character.Equipment.Head = &item
	case Tunic:
		if character.Equipment.Tunic != nil {
			character.Inventory.addItem(*character.Equipment.Tunic)
		}
		character.Equipment.Tunic = &item

	case Boots:
		if character.Equipment.Boots != nil {
			character.Inventory.addItem(*character.Equipment.Boots)
		}
		character.Equipment.Boots = &item
	}
}

func (character *Character) getHealth() int {
	return character.Health + character.Equipment.getHealthBoost()
}

func (character *Character) getMaxHealth() int {
	return character.MaxHealth + character.Equipment.getHealthBoost()
}

func (character *Character) showHealth() string {
	return fmt.Sprintf("%v/%v", character.getHealth(), character.getMaxHealth())
}

func (character *Character) spellBook(name string) {
	if Contains(make([]interface{}, len(character.Skill)), name) {
		colorFprintf("You already have the spell %v.", cyanString(name))
		return
	}
	character.Skill = append(character.Skill, "Fireball")
	character.Inventory.removeItem(name, 1)
}

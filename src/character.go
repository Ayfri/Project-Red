package main

import (
	"fmt"
	"sort"
	"strings"
)

type Character struct {
	Name       string
	Race       string
	RaceBoosts map[string]int
	Lvl        int
	MaxHealth  int
	Health     int
	Skill      []string
	Money      int
	Equipment  Equipment
	Inventory  Inventory
}

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
Race: %v
Health: %v
Lvl: %v
Money: %v
Equipment: %v
`,
		boldString(character.Name),
		greenString(character.Race),
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

func InitCharacter() {
	character = Character{
		Name:      "Ayfri",
		Race:      "Elfe",
		Lvl:       1,
		MaxHealth: 100,
		Health:    40,
		Money:     100,
		Skill: []string{
			"Punch",
		},
		Inventory: Inventory{
			"Health Potions": Item{
				Count: 3,
				Name:  "Health Potion",
				Price: 0,
				OnUse: func(item Item) {
					character.takeHealthPotion()
				},
			},
		},
	}
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

func InitInteractiveCharacter() {
	character = Character{}
	colorFprintf("What is your %v ?\n", blueString("name"))
	input := InputTextTrimmed("name")
	character.Name = strings.Title(input)
	colorFprintf("What is your Race ?\n")
	race := RaceChooser()
	character.Race = race
	character.Health = character.getMaxHealth() / 2
	character.Lvl = 1
	character.Skill = append(character.Skill, "Punch")
}

func RaceChooser() string {
	var result string
	sort.Strings(races)
	for index, race := range races {
		colorFprintf("%v. %v\n", cyanString(str(index)), race)
	}

	for {
		number, _ := InputNumber()
		if number < 0 || number > len(races) {
			continue
		}

		result = races[number]
		break
	}
	return result
}


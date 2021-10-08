package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Character struct {
	Name       string
	Race       Race
	RaceBoosts map[string]int
	Lvl        int
	MaxHealth  int
	Health     int
	Skill      []string
	Money      int
	Equipment  Equipment
	Inventory  Inventory
}

type AttackType int

const (
	Melee AttackType = iota
	Fire
	Poison
	Magic
)

func (character *Character) attack(monster *Monster) {
	weapon := character.Equipment.Weapon
	damages := weapon.AttackDamage
	if boost, ok := character.Race.Boosts["UnarmedDamages"]; ok && weapon.Name == "" {
		damages += boost
	}
	switch weapon.AttackType {
	case Melee:
		damages = weapon.AttackDamage
	case Fire:
		damages = int(float32(weapon.AttackDamage) * 0.8)
		character.parallelAttack(monster, 3, 3, weapon.AttackDamage / 10)
	case Poison:
		damages = weapon.AttackDamage / 2
		character.parallelAttack(monster, 5, 10, weapon.AttackDamage / 2)
	case Magic:
		damages = 0
		character.parallelAttack(monster, 4, 4, weapon.AttackDamage)
	}
	monster.HandleAttack(weapon, damages)
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
		greenString(character.Race.Name),
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

func (character *Character) HandleAttack(weapon *Item, damages int) {
	if boost, ok := character.Race.Boosts["MagicResistance"]; weapon.AttackType == Magic && ok {
		damages = int(float32(boost) / float32(boost / 100))
	}
	if boost, ok := character.Race.Boosts["PoisonResistance"]; weapon.AttackType == Poison && ok {
		damages = int(float32(boost) / float32(boost / 100))
	}
	if boost, ok := character.Race.Boosts["FireResistance"]; weapon.AttackType == Fire && ok {
		damages = int(float32(boost) / float32(boost / 100))
	}
	character.Health -= damages
}


func InitCharacter() {
	character = Character{
		Name:      "Ayfri",
		Race:      races[0],
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

func (character *Character) parallelAttack(monster *Monster, times int, wait int, damages int) {
	for i := 0; i < times; i++ {
		go func() {
			time.Sleep(time.Duration(wait) * time.Second)
			monster.HandleAttack(character.Equipment.Weapon, damages)
		}()
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
	character.Race = RaceChooser()
	character.Health = character.getMaxHealth() / 2
	character.Lvl = 1
	character.Skill = append(character.Skill, "Punch")
}

func RaceChooser() Race {
	var result Race
	var raceNames = RaceNames()
	sort.Strings(raceNames)
	for index, race := range RaceNames() {
		colorFprintf("%v. %v\n", cyanString(str(index)), race)
	}

	for {
		number, _ := InputNumber()
		if number < 0 || number > len(raceNames) {
			continue
		}

		result = races[number]
		break
	}
	return result
}

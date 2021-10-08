package main

import (
	"fmt"
)

type Equipment struct {
	Head   *Item
	Tunic  *Item
	Boots  *Item
	Weapon *Item
}

type EquipmentType int

const (
	Head   EquipmentType = 1 << 6
	Tunic                = 1 << 7
	Boots                = 1 << 8
	Weapon               = 1 << 9
)

func (equipment *Equipment) GetWeaponType() string {
	switch equipment.Weapon.AttackType {
	case Fire:
		return "fire"
	case Poison:
		return "poison"
	case Magic:
		return "magic"
	default:
		return "melee"
	}
}

func (equipment *Equipment) getHealthBoost() int {
	var result int
	if equipment.Head != nil {
		result += equipment.Head.EquipHealthBoost
	}
	if equipment.Tunic != nil {
		result += equipment.Tunic.EquipHealthBoost
	}
	if equipment.Boots != nil {
		result += equipment.Boots.EquipHealthBoost
	}

	return result
}

func (character *Character) printHealth() {
	colorPrintf("Character %v Health\n", redString(character.showHealth()))
}

func (equipment *Equipment) Show() string {
	head := "None"
	tunic := "None"
	boots := "None"

	if equipment.Head != nil {
		head = equipment.Head.Name
	}

	if equipment.Tunic != nil {
		tunic = equipment.Tunic.Name
	}

	if equipment.Boots != nil {
		boots = equipment.Boots.Name
	}

	return fmt.Sprintf(
		`HEAD: %v, TUNIC: %v, BOOTS: %v`,
		blueString(head),
		blueString(tunic),
		blueString(boots),
	)
}

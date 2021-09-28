package main

import (
	"fmt"
	"github.com/fatih/color"
)

type Character struct {
	Name      string
	Class     string
	Lvl       int
	MaxHealth int
	Health    int
	Skill     []string
	Money     int
	Equipment Equipment
	Inventory Inventory
}

type Equipment struct {
	Head  *Item
	Tunic *Item
	Boots *Item
}

type EquipmentType int

const (
	Head  EquipmentType = 1 << 6
	Tunic               = 1 << 7
	Boots               = 1 << 8
)

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
	colorFprintf("Character %v Health\n", color.RedString(character.showHealth()))
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
		color.BlueString(head),
		color.BlueString(tunic),
		color.BlueString(boots),
	)
}

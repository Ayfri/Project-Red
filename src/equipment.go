package main

import "fmt"

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
	head  *Item
	tunic *Item
	boots *Item
}

type EquipmentType int

const (
	Head EquipmentType = iota
	Tunic
	Boots
)

func (equipment *Equipment) Show() string {
	head := "None"
	tunic := "None"
	boots := "None"

	if equipment.head != nil {
		head = equipment.head.name
	}

	if equipment.tunic != nil {
		tunic = equipment.tunic.name
	}

	if equipment.boots != nil {
		boots = equipment.boots.name
	}

	return fmt.Sprintf(
		`HEAD: %v, TUNIC: %v, BOOTS: %v`,
		head,
		tunic,
		boots,
	)
}

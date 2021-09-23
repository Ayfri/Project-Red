package main

import "fmt"

type Inventory map[string]int

func accessInventory(inventory Inventory, selectItem bool) {
	fmt.Println("Inventory: ")
	i := 0
	for k, v := range inventory {
		i++
		if selectItem {
			fmt.Printf("%v. ", i)
		}
		fmt.Printf("%v: %v\n", k, v)
	}
}

func takePot(name string) {
	for potName, potCount := range character.inventory {
		if name == potName {
			potCount--
		}

		if name == "Potions" {
			character.health += 50
			if character.health > character.maxHealth {
				character.health = character.maxHealth
			}
		}
	}

	fmt.Printf("Health : %v/%v", character.health, character.maxHealth)
}

func (inventory Inventory) addItem(item string, count int) {
	inventory[item] += count
}

func (inventory Inventory) removeItem(item string, count int) {
	if _, ok := inventory[item]; !ok {
		return
	}

	inventory[item] -= count
	if inventory[item] <= 0 {
		delete(inventory, item)
	}
}

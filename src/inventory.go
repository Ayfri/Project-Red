package main

import "fmt"

type Item struct{
	count int
	name string
	price int
}
type Inventory map[string]Item

func accessInventory(inventory Inventory) {
	fmt.Println("Inventory: ")
	i := 0
	for name, item := range inventory {
		i++
		fmt.Printf("%v. %v: %v\n", i, name, item)
	}
}

func accessPricedInventory(inventory Inventory) {
	i := 0
	for name, item := range inventory {
		i++
		fmt.Printf("%v. %v: %v (Price: %v)\n", i, name, item.count, item.price)
	}
}

func takePot(name string) {
	for potName := range character.inventory {
		if name == potName {
			character.inventory.removeItem(name, 1)
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

func (inventory *Inventory) addItem(item Item) {
	if val, ok := (*inventory)[item.name]; ok {
		val.count += item.count
		(*inventory)[item.name] = val
	} else {
		(*inventory)[item.name] = item
	}
}

func (inventory *Inventory) removeItem(name string, count int) {
	if item, ok := (*inventory)[name]; !ok {
		return
	} else {
		item.count -= count
		(*inventory)[name] = item
		if item.count <= 0 {
			delete(*inventory, name)
		}
	}
}

func (inventory *Inventory) debug() string {
	var result string

	for name, item := range *inventory {
		result += fmt.Sprintf("{%v, count=%d, name=%s, price=%d}", name, item.count, item.name, item.price)
	}

	return result
}

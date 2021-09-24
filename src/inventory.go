package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Inventory map[string]Item

type SelectorType int

const (
	Merchant SelectorType = iota
	PlayerInventory
	Blacksmith
)

func (inventory Inventory) show(selectorType SelectorType) {
	keys := inventory.keys()
	sort.Strings(keys)

	for i, name := range keys {
		item := inventory[name]
		switch selectorType {
		case Merchant:
			fmt.Printf("%v. %v: %v (Price: %v)\n", i+1, name, item.count, item.price)
		case PlayerInventory:
			fmt.Printf("%v. %v: %v\n", i+1, name, item.count)
		case Blacksmith:
			fmt.Printf("%v. %v (Requires: %v)\n", i+1, name, item.forgingRequires)
		}
	}
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

func (inventory *Inventory) makeSelector(selectorType SelectorType, whenQuit func()) {
	fmt.Println("Select the item you want : (q to quit)")
	keys := inventory.keys()
	sort.Strings(keys)
	inventory.show(selectorType)

	for {
		input, _ := reader.ReadString('\n')
		number, err := strconv.Atoi(strings.TrimSpace(input))
		if input[0] == 'q' {
			whenQuit()
			break
		}

		if err != nil || number < 0 || number > len(merchant) {
			continue
		}

		i := 1
		for _, name := range keys {
			item := (*inventory)[name]
			if number == i {
				switch selectorType {
				case Merchant:
					if character.money < item.price {
						fmt.Printf("You need '%v' more money to buy '%v'.", -(character.money - item.price), item.name)
						continue
					}
					inventory.removeItem(name, i)

					receivingItem := item
					receivingItem.count = 1
					character.inventory.addItem(receivingItem)
					fmt.Printf("One '%v' bought.\n", item.name)
				case PlayerInventory:
					if item.onUse != nil {
						item.onUse(item)
					}
					inventory.removeItem(name, i)
					if item.equipmentType == Head || item.equipmentType == Tunic || item.equipmentType == Boots {
						fmt.Printf("'%v' equiped.\n", item.name)
					} else {
						fmt.Printf("One '%v' used.\n", item.name)
					}
				case Blacksmith:
					if canForge, forgeErr := character.canForge(item); canForge {
						character.forgeItem(item)
						inventory.removeItem(name, i)
						fmt.Printf("One '%v' crafted.\n", item.name)
					} else {
						fmt.Println(forgeErr)
					}
				}

				break
			}
			i++
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

package main

import (
	"fmt"
	"strings"
)

var blacksmith = Inventory{
	"Adventurer's Hats": Item{
		count:           1,
		name:            "Adventurer's Hat",
		forgingRequires: map[string]int{crowFeather: 1, boarFur: 1},
		forgingPrice:    5,
		equipmentType: 	Head,
		onUse: func(item Item) {
			character.equip(item)
		},
	},
	"Adventurer's Tunics": Item{
		count:           1,
		name:            "Adventurer's Tunic",
		forgingRequires: map[string]int{wolfFur: 2, boarFur: 1},
		forgingPrice:    5,
		equipmentType: Tunic,
		onUse: func(item Item) {
			character.equip(item)
		},
	},
	"Adventurer's Boots": Item{
		count:           1,
		name:            "Adventurer's Boots",
		forgingRequires: map[string]int{wolfFur: 1, boarFur: 1},
		forgingPrice:    5,
		equipmentType: Boots,
		onUse: func(item Item) {
			character.equip(item)
		},
	},
}

func (character *Character) canForge(item Item) (bool, string) {
	if character.money-item.forgingPrice < 0 {
		return false, fmt.Sprintf("You don't have enough money to forge '%v'.", item.name)
	}

	missingItems := DifferentKeys(item.forgingRequires, character.inventory)

	if len(missingItems) > 0 {
		return false, fmt.Sprintf("You need '%v' to craft '%v'.", strings.Join(missingItems, " & "), item.name)
	}

	if item.forgingRequires != nil {
		for _, inventoryItem := range character.inventory {
			for name, count := range item.forgingRequires {
				if strings.Contains(inventoryItem.name, name) && inventoryItem.count-count < 0 {
					return false, fmt.Sprintf("You need %v more '%v' to craft this item.\n", -(inventoryItem.count - count), inventoryItem.name)
				}
			}
		}
	}

	return true, ""
}

func (character *Character) forgeItem(item Item) {
	if canForge, err := character.canForge(item); !canForge {
		fmt.Println(err)
		return
	}

	if item.forgingRequires != nil {
		for _, inventoryItem := range character.inventory {
			for name, count := range item.forgingRequires {
				if strings.Contains(inventoryItem.name, name) {
					inventoryItem.count -= count
				}
			}
		}
	}

	resultingItem := item
	resultingItem.count = 1
	character.inventory.addItem(resultingItem)
}

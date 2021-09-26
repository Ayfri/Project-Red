package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type ForgingRequires map[string]int

var blacksmith = Inventory{
	"Adventurer's Hats": Item{
		count:           1,
		name:            "Adventurer's Hat",
		forgingRequires: ForgingRequires{crowFeather: 1, boarFur: 1},
		forgingPrice:    5,
		equipmentType:   Head,
		onUse: func(item Item) {
			character.equip(item)
		},
	},
	"Adventurer's Tunics": Item{
		count:           1,
		name:            "Adventurer's Tunic",
		forgingRequires: ForgingRequires{wolfFur: 2, boarFur: 1},
		forgingPrice:    5,
		equipmentType:   Tunic,
		onUse: func(item Item) {
			character.equip(item)
		},
	},
	"Adventurer's Boots": Item{
		count:           1,
		name:            "Adventurer's Boots",
		forgingRequires: ForgingRequires{wolfFur: 1, boarFur: 1},
		forgingPrice:    5,
		equipmentType:   Boots,
		onUse: func(item Item) {
			character.equip(item)
		},
	},
}

func (character *Character) canForge(item Item) (bool, string) {
	if character.money-item.forgingPrice < 0 {
		return false, fmt.Sprintf("You don't have enough money to forge %v.", color.BlueString(item.name))
	}

	missingItems := DifferentKeys(item.forgingRequires, character.inventory)

	if len(missingItems) > 0 {
		return false, fmt.Sprintf(
			"You need %v to craft %v.",
			color.BlueString(strings.Join(missingItems, " & ")),
			color.BlueString(item.name),
		)
	}

	if item.forgingRequires != nil {
		for _, inventoryItem := range character.inventory {
			for name, count := range item.forgingRequires {
				if strings.Contains(inventoryItem.name, name) && inventoryItem.count-count < 0 {
					return false, fmt.Sprintf(
						"You need %v more %v to craft this item.\n",
						color.CyanString(str(-(inventoryItem.count - count))),
						color.BlueString(inventoryItem.name),
					)
				}
			}
		}
	}

	return true, ""
}

func (character *Character) forgeItem(item Item) {
	if canForge, err := character.canForge(item); !canForge {
		colorFprintf(err)
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

// TODO : Delete this func
func (forgingRequires *ForgingRequires) show() string {
	var result []string
	for name, count := range *forgingRequires {
		result = append(result, fmt.Sprintf("%v:%v", color.BlueString(name), color.CyanString(str(count))))
	}
	return strings.Join(result, ", ")
}
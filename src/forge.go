package main

import (
	"fmt"
	"strings"
)

type ForgingRequires map[string]int

var blacksmith = NPC{
	Inventory: Inventory{
		"Adventurer's Hats": Item{
			Count:            1,
			Name:             "Adventurer's Hat",
			ForgingRequires:  ForgingRequires{crowFeather: 1, boarFur: 1},
			ForgingPrice:     5,
			EquipmentType:    Head,
			EquipHealthBoost: 10,
			OnUse: func(item Item) {
				character.equip(item)
			},
		},
		"Adventurer's Tunics": Item{
			Count:            1,
			Name:             "Adventurer's Tunic",
			ForgingRequires:  ForgingRequires{wolfFur: 2, boarFur: 1},
			ForgingPrice:     5,
			EquipmentType:    Tunic,
			EquipHealthBoost: 25,
			OnUse: func(item Item) {
				character.equip(item)
			},
		},
		"Adventurer's Boots": Item{
			Count:            1,
			Name:             "Adventurer's Boots",
			ForgingRequires:  ForgingRequires{wolfFur: 1, boarFur: 1},
			ForgingPrice:     5,
			EquipmentType:    Boots,
			EquipHealthBoost: 15,
			OnUse: func(item Item) {
				character.equip(item)
			},
		},
	},
	Name: "Alvor",
	Type: Blacksmith,
}

func (character *Character) canForge(item Item) (bool, string) {
	if character.Money-item.ForgingPrice < 0 {
		return false, fmt.Sprintf("You don't have enough Money to forge %v.", blueString(item.Name))
	}

	missingItems := DifferentKeys(item.ForgingRequires, character.Inventory)

	if len(missingItems) > 0 {
		return false, fmt.Sprintf(
			"You need %v to craft %v.",
			blueString(strings.Join(missingItems, " & ")),
			blueString(item.Name),
		)
	}

	if item.ForgingRequires != nil {
		for _, inventoryItem := range character.Inventory {
			for name, count := range item.ForgingRequires {
				if strings.Contains(inventoryItem.Name, name) && inventoryItem.Count-count < 0 {
					return false, fmt.Sprintf(
						"You need %v more %v to craft this item.\n",
						cyanString(str(-(inventoryItem.Count - count))),
						blueString(inventoryItem.Name),
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

	if item.ForgingRequires != nil {
		for _, inventoryItem := range character.Inventory {
			for name, count := range item.ForgingRequires {
				if strings.Contains(inventoryItem.Name, name) {
					inventoryItem.Count -= count
				}
			}
		}
	}

	resultingItem := item
	resultingItem.Count = 1
	character.Inventory.addItem(resultingItem)
}

// TODO : Delete this func
// From stream twitch.tv/ayfri1015, from ArFTNL.
func (forgingRequires *ForgingRequires) show() string {
	var result []string
	for name, count := range *forgingRequires {
		result = append(result, fmt.Sprintf("%v:%v", blueString(name), cyanString(str(count))))
	}
	return strings.Join(result, ", ")
}

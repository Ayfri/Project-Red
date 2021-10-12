package main

import (
	"fmt"
	"strings"
)

type ForgingRequires map[string]int

var blacksmith = NPC{
	Inventory: Inventory{
		adventurerHat: Item{
			Count:            1,
			Name:             adventurerHat,
			ForgingRequires:  ForgingRequires{fiber: 1, boarFur: 1},
			ForgingPrice:     5,
			EquipmentType:    Head,
			EquipHealthBoost: 10,
			OnUse: func(item Item) {
				character.equip(item)
			},
		},
		adventurerTunic: Item{
			Count:            1,
			Name:             adventurerTunic,
			ForgingRequires:  ForgingRequires{iron: 2, boarFur: 1},
			ForgingPrice:     5,
			EquipmentType:    Tunic,
			EquipHealthBoost: 25,
			OnUse: func(item Item) {
				character.equip(item)
			},
		},
		adventurerBoot: Item{
			Count:            1,
			Name:             adventurerBoot,
			ForgingRequires:  ForgingRequires{iron: 1, boarFur: 1},
			ForgingPrice:     5,
			EquipmentType:    Boots,
			EquipHealthBoost: 15,
			OnUse: func(item Item) {
				character.equip(item)
			},
		},
		dragonBane: Item{
			Count:            1,
			Name:             dragonBane,
			ForgingRequires:  ForgingRequires{iron: 4, fiber: 1},
			ForgingPrice:     5,
			EquipmentType:    Weapon,
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
		colorPrintf(err)
		return
	}

	if !character.canAddItem(item) {
		yellow("You can't add any new item to your inventory !\n")
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

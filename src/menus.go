package main

import (
	"fmt"
	"github.com/fatih/color"
	"sort"
	"strconv"
	"strings"
)

func combatMenu(character *Character, enemy *Monster) bool {
	fmt.Println("1. Attack")
	fmt.Println("2. Inventory")

	input, _ := reader.ReadString('\n')
	number, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil || number < 0 || number > len(merchant) {
		return false
	}

	if input[0] == 'q' {
		return true
	}

	switch number {
	case 1:
		character.attack(enemy)
	case 2:
		character.Inventory.makeSelector(PlayerInventory, func(){})
	}
	return false
}

func (inventory *Inventory) makeSelector(selectorType SelectorType, whenQuit func()) {
	colorFprintf(boldString("Select the item you want : %v\n"), color.RedString("(q to quit)"))
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
					if character.Money < item.Price {
						colorFprintf(
							"You need %v more money to buy %v.\n",
							color.YellowString(str(-(character.Money - item.Price))),
							color.BlueString(item.Name),
						)
						break
					}
					inventory.removeItem(name, 1)

					receivingItem := item
					receivingItem.Count = 1
					character.Money -= item.Price
					character.Inventory.addItem(receivingItem)
					printItemTaken("One %v bought.\n", item.Name)
				case PlayerInventory:
					if item.OnUse != nil {
						item.OnUse(item)
					}
					inventory.removeItem(name, 1)
					if item.EquipmentType == Head || item.EquipmentType == Tunic || item.EquipmentType == Boots {
						printItemTaken("%v equipped.\n", item.Name)
					} else {
						printItemTaken("One %v used.\n", item.Name)
					}
				case Blacksmith:
					if canForge, forgeErr := character.canForge(item); canForge {
						character.forgeItem(item)
						inventory.removeItem(name, 1)
						printItemTaken("One %v crafted.\n", item.Name)
					} else {
						fmt.Println(forgeErr)
					}
				}
				keys = inventory.keys()
				sort.Strings(keys)
				break
			}
			i++
		}
	}
}

func showMainMenu() {
	boldFunc("Select something :")
	color.Cyan("1: Show character information.")
	color.Green("2: Show Inventory.")
	color.Blue("3: Speak to Merchant.")
	color.Magenta("4: Speak to Blacksmith.")
	color.Yellow("5: Fight with training goblin.")
	color.Red("6: Quit")
}

package main

import (
	"fmt"
	"sort"
)

func combatMenu(turn int, character *Character, enemy *Monster) bool {
	printCenteredTitle(fmt.Sprintf("Turn %d", turn))
	fmt.Println("1. Attack")
	fmt.Println("2. Inventory")

	number, quit := InputNumber()

	if number <= 0 || number > len(merchant) {
		return false
	}

	if quit {
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
	colorFprintf(boldString("Select the item you want : %v\n", redString("(q to quit)")))
	keys := inventory.keys()
	sort.Strings(keys)
	inventory.show(selectorType)

	for {
		number, quit := InputNumber()

		if quit {
			whenQuit()
			break
		}

		if number <= 0 || number > len(merchant) {
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
							yellowString(str(-(character.Money - item.Price))),
							blueString(item.Name),
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
	printLine()
	bold("Select something :")
	cyan("1: Show character information.")
	green("2: Show Inventory.")
	blue("3: Speak to Merchant.")
	magenta("4: Speak to Blacksmith.")
	yellow("5: Fight with training goblin.")
	red("6: Quit")
}

package main

import (
	"fmt"
	"sort"
)

func combatMenu(turn int, character *Character, enemy *Monster) bool {
	printCenteredTitle(fmt.Sprintf("Turn %d", turn))
	colorPrintf("%v. Attack\n", cyanString(str(1)))
	colorPrintf("%v. Inventory\n", cyanString(str(2)))

	number, quit := InputNumber()

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
	colorPrintf(boldString("Select the item you want : %v\n", redString("(q to quit)")))
	keys := inventory.keys()
	sort.Strings(keys)
	inventory.show(selectorType)

	for {
		number, quit := InputNumber()

		if quit {
			whenQuit()
			return
		}

		if number <= 0 || number > len(*inventory) {
			continue
		}

		i := 1
		for _, name := range keys {
			item := (*inventory)[name]
			if number == i {
				switch selectorType {
				case MerchantInventory:
					if character.Money < item.Price {
						colorPrintf(
							"You need %v more money to buy %v.\n",
							yellowString(str(-(character.Money - item.Price))),
							blueString(item.Name),
						)
						break
					}

					if !character.canAddItem(item) {
						yellow("You can't add any new item to your inventory !\n")
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
				case PlayerSellInventory:
					inventory.removeItem(name, 1)
					character.Money += item.Price
					colorPrintf("You sold %v for %v money.", name, yellowString(str(item.Price)))
				case BlacksmithInventory:
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
	monsterText := "Fight with training goblin."
	if !isThereMonster {
		monsterText = "Search monster."
	}
	yellow("5: %v", monsterText)
	red("6: Quit")
}

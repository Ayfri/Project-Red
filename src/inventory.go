package main

import (
	"fmt"
	"github.com/fatih/color"
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

func (inventory *Inventory) show(selectorType SelectorType) {
	keys := inventory.keys()
	sort.Strings(keys)

	for i, name := range keys {
		item := (*inventory)[name]
		count := str(item.count)
		index := color.CyanString(str(i + 1))

		switch selectorType {
		case Merchant:
			colorFprintf("%v. %v: %v %v\n", index, name, color.GreenString(count), color.YellowString("(Price: %v)", str(item.price)))
		case PlayerInventory:
			colorFprintf("%v. %v: %v\n", index, name, color.GreenString(count))
		case Blacksmith:
			colorFprintf("%v. %v (Requires: %v)\n", index, name, item.forgingRequires.show())
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
					if character.money < item.price {
						colorFprintf(
							"You need %v more money to buy %v.\n",
							color.YellowString(str(-(character.money - item.price))),
							color.BlueString(item.name),
						)
						continue
					}
					inventory.removeItem(name, i)

					receivingItem := item
					receivingItem.count = 1
					character.money -= item.price
					character.inventory.addItem(receivingItem)
					itemTaken("One %v bought.\n", item.name)
				case PlayerInventory:
					if item.onUse != nil {
						item.onUse(item)
					}
					inventory.removeItem(name, i)
					if item.equipmentType == Head || item.equipmentType == Tunic || item.equipmentType == Boots {
						itemTaken("%v equipped.\n", item.name)
					} else {
						itemTaken("One %v used.\n", item.name)
					}
				case Blacksmith:
					if canForge, forgeErr := character.canForge(item); canForge {
						character.forgeItem(item)
						inventory.removeItem(name, i)
						itemTaken("One %v crafted.\n", item.name)
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

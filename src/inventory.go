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
		count := str(item.Count)
		index := color.CyanString(str(i + 1))

		switch selectorType {
		case Merchant:
			colorFprintf("%v. %v: %v %v\n", index, name, color.GreenString(count), color.YellowString("(Price: %v)", str(item.Price)))
		case PlayerInventory:
			colorFprintf("%v. %v: %v\n", index, name, color.GreenString(count))
		case Blacksmith:
			colorFprintf("%v. %v (Requires: %v)\n", index, name, item.ForgingRequires.show())
		}
	}
}

func (inventory *Inventory) addItem(item Item) {
	if val, ok := (*inventory)[item.Name]; ok {
		val.Count += item.Count
		(*inventory)[item.Name] = val
	} else {
		(*inventory)[item.Name] = item
	}
}

func (inventory *Inventory) removeItem(name string, count int) {
	if item, ok := (*inventory)[name]; !ok {
		return
	} else {
		item.Count -= count
		(*inventory)[name] = item
		if item.Count <= 0 {
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
					if character.Money < item.Price {
						colorFprintf(
							"You need %v more money to buy %v.\n",
							color.YellowString(str(-(character.Money - item.Price))),
							color.BlueString(item.Name),
						)
						break
					}
					inventory.removeItem(name, i)

					receivingItem := item
					receivingItem.Count = 1
					character.Money -= item.Price
					character.Inventory.addItem(receivingItem)
					itemTaken("One %v bought.\n", item.Name)
				case PlayerInventory:
					if item.OnUse != nil {
						item.OnUse(item)
					}
					inventory.removeItem(name, i)
					if item.EquipmentType == Head || item.EquipmentType == Tunic || item.EquipmentType == Boots {
						itemTaken("%v equipped.\n", item.Name)
					} else {
						itemTaken("One %v used.\n", item.Name)
					}
				case Blacksmith:
					if canForge, forgeErr := character.canForge(item); canForge {
						character.forgeItem(item)
						inventory.removeItem(name, i)
						itemTaken("One %v crafted.\n", item.Name)
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
		result += fmt.Sprintf("{%v, Count=%d, Name=%s, Price=%d}", name, item.Count, item.Name, item.Price)
	}

	return result
}

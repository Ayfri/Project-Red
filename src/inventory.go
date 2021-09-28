package main

import (
	"github.com/fatih/color"
	"sort"
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

func (inventory *Inventory) debug() {
	for _, item := range *inventory {
		colorFprintf("{Count=%s, Name=%s, Price=%s}\n", color.CyanString(str(item.Count)), color.BlueString(item.Name), color.YellowString(str(item.Price)))
	}
}

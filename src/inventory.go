package main

import (
	"sort"
)

type Inventory map[string]Item

type SelectorType int

const (
	MerchantInventory SelectorType = iota
	PlayerInventory
	PlayerSellInventory
	BlacksmithInventory
)

func (inventory *Inventory) show(selectorType SelectorType) {
	keys := inventory.keys()
	sort.Strings(keys)

	for i, name := range keys {
		item := (*inventory)[name]
		count := str(item.Count)
		index := cyanString(str(i + 1))

		switch selectorType {
		case MerchantInventory:
			colorPrintf("%v. %v: %v %v\n", index, name, greenString(count), yellowString("(Price: %v)", str(item.Price)))
		case PlayerInventory:
			colorPrintf("%v. %v: %v\n", index, name, greenString(count))
		case BlacksmithInventory:
			colorPrintf("%v. %v (Requires: %v)\n", index, name, item.ForgingRequires.show())
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

func (inventory *Inventory) canAddItem(item Item, limit int) bool {
	return inventory.getInventoryLength()+item.Count <= limit
}

func (inventory *Inventory) getInventoryLength() int {
	var result int
	for _, item := range *inventory {
		result += item.Count
	}
	return result
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
		colorPrintf("{Count=%s, Name=%s, Price=%s}\n", cyanString(str(item.Count)), blueString(item.Name), yellowString(str(item.Price)))
	}
}

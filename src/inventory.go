package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Inventory map[string]Item

func (inventory Inventory) show(priced bool) {
	i := 0
	for name, item := range inventory {
		i++
		if priced {
			fmt.Printf("%v. %v: %v (Price: %v)\n", i, name, item.count, item.price)
		} else {
			fmt.Printf("%v. %v: %v\n", i, name, item.count)
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

func (inventory *Inventory) makeSelector(priced bool, whenQuit func()) {
	fmt.Println("Select the item you want : (q to quit)")
	inventory.show(priced)

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
		for name, item := range *inventory {
			if number == i {
				inventory.removeItem(name, i)

				if priced {
					fmt.Printf("One '%v' bought.\n", item.name)
				} else {
					if item.onUse != nil {
						item.onUse()
					}
					fmt.Printf("One '%v' used.\n", item.name)
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

package main

import (
	"bufio"
	"fmt"
	"os"
)

var character Character
var reader = bufio.NewReader(os.Stdin)

func main() {
	Init()
	var input string
	showMainMenu()
	for {
		input, _ = reader.ReadString('\n')
		switch input[0] {
		case '1':
			character.displayInfo()
			showMainMenu()
		case '2':
			fmt.Println("Inventory:")
			character.inventory.makeSelector(PlayerInventory, showMainMenu)
		case '3':
			fmt.Println("Merchant:")
			merchant.makeSelector(Merchant, showMainMenu)
		case '4':
			fmt.Println("Blacksmith:")
			blacksmith.makeSelector(Blacksmith, showMainMenu)
		case '5':
			os.Exit(1)
		}
	}
}

func showMainMenu() {
	fmt.Println(`
Select something :

1: Show character information.
2: Show inventory.
3: Speak to Merchant.
4: Speak to blacksmith.
5: Quit`)
}

func Init() {
	character = Character{
		name:      "Ayfri",
		class:     "Elfe",
		lvl:       1,
		maxHealth: 100,
		health:    40,
		money:     100,
		skill: []string{
			"Punch",
		},
		inventory: Inventory{
			"Health Potions": Item{
				count: 3,
				name:  "Health Potion",
				price: 0,
				onUse: func() {
					character.takeHealthPotion()
				},
			},
		},
	}

	for _, item := range blacksmith {
		item.onUse = func() {
			character.forgeItem(item)
		}
	}
}

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
			character.Inventory.makeSelector(PlayerInventory, showMainMenu)
		case '3':
			fmt.Println("Merchant:")
			merchant.makeSelector(Merchant, showMainMenu)
		case '4':
			fmt.Println("Blacksmith:")
			blacksmith.makeSelector(Blacksmith, showMainMenu)
		case '5':
			trainingFight(&character, &trainingGoblin)
			showMainMenu()
		case '6':
			os.Exit(1)
		}
	}
}

func Init() {
	character = Character{
		Name:      "Ayfri",
		Class:     "Elfe",
		Lvl:       1,
		MaxHealth: 100,
		Health:    40,
		Money:     100,
		Skill: []string{
			"Punch",
		},
		Inventory: Inventory{
			"Health Potions": Item{
				Count: 3,
				Name:  "Health Potion",
				Price: 0,
				OnUse: func(item Item) {
					character.takeHealthPotion()
				},
			},
		},
	}
	trainingGoblin = InitGoblin("Training Goblin", 40, 5)
}

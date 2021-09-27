package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
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
			os.Exit(1)
		}
	}
}

func showMainMenu() {
	boldFunc("Select something :")
	color.Cyan("1: Show character information.")
	color.Green("2: Show Inventory.")
	color.Magenta("3: Speak to Merchant.")
	color.Yellow("4: Speak to Blacksmith.")
	color.Red("5: Quit")
}

func Init() {
	character = Character{
		Name:      "Ayfri",
		Class:     "Elfe",
		Lvl:       1,
		MaxHealth: 100,
		Health:    40,
		Money:     1,
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
}

package main

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var character Character

func main() {
	InitColors()
	InitInteractiveCharacter()
	trainingGoblin = InitGoblin("Training Goblin", 40, 5)
	showMainMenu()
	for {
		input, quit := InputNumber()
		if quit {
			os.Exit(1)
		}

		switch input {
		case 1:
			character.displayInfo()
			showMainMenu()
		case 2:
			printCenteredTitle("Inventory")
			character.Inventory.makeSelector(PlayerInventory, showMainMenu)
		case 3:
			printCenteredTitle("Merchant")
			merchant.makeSelector(Merchant, showMainMenu)
		case 4:
			printCenteredTitle("Blacksmith")
			blacksmith.makeSelector(Blacksmith, showMainMenu)
		case 5:
			trainingFight(&character, &trainingGoblin)
			showMainMenu()
		case 6:
			os.Exit(1)
		}
	}
}

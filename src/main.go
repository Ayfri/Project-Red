package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var reader = bufio.NewReader(os.Stdin)
var character Character

func main() {
	rand.Seed(time.Now().UnixNano())
	InitColors()
	InitInteractiveCharacter()
	monster = InitGoblin("Training Goblin", 40, 5)
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
			merchant.speak()
			showMainMenu()
		case 4:
			printCenteredTitle("Blacksmith")
			blacksmith.speak()
			showMainMenu()
		case 5:
			trainingFight(&character, &monster)
			showMainMenu()
		case 6:
			os.Exit(1)
		case 9:
			character.Money = 10000
		}
	}
}

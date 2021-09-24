package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
			displayInfo(character)
			showMainMenu()
		case '2':
			accessInventory(character.inventory)
			fmt.Println("Press q to quit.")
			for {
				input, _ = reader.ReadString('\n')
				number, err := strconv.Atoi(strings.TrimSpace(input))
				if input[0] == 'q' {
					break
				}

				if err != nil || number < 0 || number > len(merchant) {
					continue
				}

				i := 1
				for name := range character.inventory {
					if number == i {
						character.inventory.removeItem(name, 1)
						fmt.Printf("One '%v' used.\n", name)
						if strings.Contains(strings.ToLower(name), "spellbook") {
							character.spellBook(name)
						}
						break
					}
					i++
				}
			}
		case '3':
			displayMerchant()
			showMainMenu()
		case '4':
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
4: Quit`)
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
			},
		},
	}
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var merchant = Inventory{
	"Health Potions": Item{
		count: 1,
		name:  "Health Potion",
		price: 3,
	},
	"Poisonous Potion": Item{
		count: 1,
		name:  "Poisonous Potion",
		price: 6,
	},
	"Spellbook : Fireball": Item{
		count: 1,
		name:  "Spellbook : Fireball",
		price: 25,
	},
}

func displayMerchant() {
	fmt.Println("Select the item you want : (q to quit)")
	accessPricedInventory(merchant)
	for {
		input, _ := reader.ReadString('\n')
		number, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil || number < 0 || number > len(merchant) {
			continue
		}

		if input[0] == 'q' {
			showMainMenu()
			break
		}

		i := 1
		for name, vendingItem := range merchant {
			if number == i {
				item := vendingItem
				item.count = 1
				merchant.removeItem(name, 1)
				character.inventory.addItem(item)
				fmt.Printf("One '%v' bought.\n", name)
			}
			i++
		}
	}
}

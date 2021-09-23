package main

import (
	"fmt"
	"strconv"
	"strings"
)

var merchant = Inventory{
	"Potions": 1,
	"Spellbook : Fireball": 1,

}

func displayMerchant() {
	fmt.Println("Merchant.")
	accessInventory(merchant, true)
	for {
		input, _ := reader.ReadString('\n')
		number, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil || number < 0 || number > len(merchant) {
			continue
		}

		i := 1
		for name := range merchant {
			if number == i {
				merchant.removeItem(name, 1)
				character.inventory.addItem(name, 1)
				fmt.Printf("One '%v' bought.\n", name)
				break
			}
			i++
		}
		break
	}
}

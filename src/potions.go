package main

import (
	"fmt"
	"time"
)

func (character *Character) takePoisonPotion() {
	for _, item := range character.Inventory {
		if item.Name == healthPotion {
			for i := 0; i < 3; i++ {
				character.Health -= 10
				time.Sleep(time.Second)
				fmt.Printf("You took 10 damages.\n")
				character.showHealth()
			}
			character.Inventory.removeItem(poisonousPotion, 1)
		}
	}
}

func (character *Character) takeHealthPotion() {
	for _, item := range character.Inventory {
		if item.Name == healthPotion {
			character.Health += 50
			if character.Health > character.MaxHealth {
				character.Health = character.MaxHealth
			}
			character.Inventory.removeItem(healthPotion, 1)
		}
	}

	character.showHealth()
}

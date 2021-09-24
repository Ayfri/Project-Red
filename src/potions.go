package main

import (
	"fmt"
	"time"
)

func (character *Character) takePoisonPotion() {
	for _, item := range character.inventory {
		if item.name == healthPotion {
			for i := 0; i < 3; i++ {
				character.health -= 10
				time.Sleep(time.Second)
				fmt.Printf("You took 10 damages.\n")
				character.showHealth()
			}
			character.inventory.removeItem(poisonousPotion, 1)
		}
	}
}

func (character * Character) takeHealthPotion() {
	for _, item := range character.inventory {
		if item.name == healthPotion {
			character.health += 50
			if character.health > character.maxHealth {
				character.health = character.maxHealth
			}
			character.inventory.removeItem(healthPotion, 1)
		}
	}

	character.showHealth()
}

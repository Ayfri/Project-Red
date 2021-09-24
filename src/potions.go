package main

import (
	"fmt"
	"time"
)

func (character *Character) takePoisonPotion() {
	for i := 0; i < 3; i++ {
		character.health -= 10
		time.Sleep(time.Second)
		fmt.Printf("You took 10 damages.\n")
		character.showHealth()
	}
}

func (character * Character) takeHealthPotion() {
	for _, item := range character.inventory {
		if item.name == "Health Potion" {
			character.health += 50
			if character.health > character.maxHealth {
				character.health = character.maxHealth
			}
		}
	}

	character.showHealth()
}

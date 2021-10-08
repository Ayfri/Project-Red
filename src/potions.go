package main

import (
	"fmt"
	"time"
)

func (character *Character) takePoisonPotion() {
	for _, item := range character.Inventory {
		if item.Name == poisonousPotion {
			for i := 0; i < 3; i++ {
				character.Health -= 10
				time.Sleep(time.Second)
				fmt.Printf("You took 10 damages.\n")
				character.printHealth()
			}
			character.Inventory.removeItem(poisonousPotion, 1)
		}
	}
}

func (character *Character) throwPoisonPotion(monster *Monster) {
	for _, item := range character.Inventory {
		if item.Name == poisonousPotion {
			colorPrintf("You throw one poison potion to %v.\n", redString(monster.Name))
			character.Inventory.removeItem(poisonousPotion, 1)
			damages := 5
			for i := 0; i < 3; i++ {
				go func() {
					if monster.Health < 0 {
						return
					}
					time.Sleep(time.Duration(5) * time.Second)
					monster.HandleAttack(&Item{AttackType: Poison}, damages)
					colorPrintf("%v took poison damage, %v damages taken.\n", redString(monster.Name), redString(str(damages)))
					monster.showHealth()
				}()
			}
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

	character.printHealth()
}

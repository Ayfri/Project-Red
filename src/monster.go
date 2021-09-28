package main

import (
	"fmt"
	"github.com/fatih/color"
)

type Monster struct {
	Name         string
	MaxHealth    int
	Health       int
	AttackDamage int
}

var trainingGoblin Monster

func InitGoblin(name string, maxHealth int, attackDamage int) Monster {
	return Monster {
	    Name: name,
		MaxHealth: maxHealth,
		Health: maxHealth,
		AttackDamage: attackDamage,
	}
}

func (monster *Monster) attack(character *Character) {
	character.Health -= monster.AttackDamage
	printAttack(*monster, *character, monster.AttackDamage)
}

func (monster *Monster) goblinPattern(turn int, character *Character) {
	if turn % 3 == 0 {
		monster.specialAttack(character)
	} else {
		monster.attack(character)
	}
}

func (monster *Monster) printHealth() {
	colorFprintf("Monster %v Health\n", color.RedString(monster.showHealth()))
}

func (monster *Monster) showHealth() string {
	return fmt.Sprintf("%v/%v", monster.Health, monster.MaxHealth)
}


func (monster *Monster) specialAttack(character *Character) {
	damages := monster.AttackDamage * 2
	character.Health -= damages
	printAttack(*monster, *character, damages)
}

func trainingFight(character *Character, monster *Monster) {
	turn := 0

	color.Yellow("Press q to abandon")
	for {
		turn++
		stop := combatMenu(character, monster)
		if stop {
			break
		}
		monster.goblinPattern(turn, character)
		if character.getHealth() <= 0 {
			character.dead()
			break
		}

		if monster.Health <= 0 {
			colorFprintf("Monster %v dead, you won !\n", color.BlueString(monster.Name))
			break
		}
	}
	colorFprintf("Combat terminated in %v turn.\n", color.CyanString(str(turn)))
}

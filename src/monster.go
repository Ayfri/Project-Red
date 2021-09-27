package main

import "fmt"

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

	for {
		turn++
		monster.attack(character)
		character.attack(monster)
		if character.getHealth() <= 0 {
			character.dead()
		}
	}
}

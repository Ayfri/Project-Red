package main

import (
	"fmt"
)

type Monster struct {
	AttackDamage int
	Health       int
	MaxHealth    int
	Name         string
	Race         Race
}

var trainingGoblin Monster

func InitGoblin(name string, maxHealth int, attackDamage int) Monster {
	return Monster{
		Name:         name,
		MaxHealth:    maxHealth,
		Health:       maxHealth,
		AttackDamage: attackDamage,
	}
}

func (monster *Monster) attack(character *Character) {
	character.Health -= monster.AttackDamage
	printAttack(*monster, *character, monster.AttackDamage)
}

func (monster *Monster) goblinPattern(turn int, character *Character) {
	if turn%3 == 0 {
		monster.specialAttack(character)
	} else {
		monster.attack(character)
	}
}

func (monster *Monster) printHealth() {
	colorFprintf("Monster %v Health\n", redString(monster.showHealth()))
}

func (monster *Monster) showHealth() string {
	return fmt.Sprintf("%v/%v", monster.Health, monster.MaxHealth)
}

func (monster *Monster) specialAttack(character *Character) {
	damages := monster.AttackDamage * 2
	character.Health -= damages
	printAttack(*monster, *character, damages)
}

func (monster *Monster) HandleAttack(weapon *Item, damages int) {
	if boost, ok := monster.Race.Boosts["MagicResistance"]; weapon.AttackType == Magic && ok {
		damages = int(float32(boost) / float32(boost / 100))
	}
	if boost, ok := monster.Race.Boosts["PoisonResistance"]; weapon.AttackType == Poison && ok {
		damages = int(float32(boost) / float32(boost / 100))
	}
	if boost, ok := monster.Race.Boosts["FireResistance"]; weapon.AttackType == Fire && ok {
		damages = int(float32(boost) / float32(boost / 100))
	}
	monster.Health -= damages
}

func trainingFight(character *Character, monster *Monster) {
	turn := 0

	yellow("Press q to abandon")
	for {
		turn++
		stop := combatMenu(turn, character, monster)
		if stop {
			break
		}
		monster.goblinPattern(turn, character)
		if character.getHealth() <= 0 {
			character.dead()
			break
		}

		if monster.Health <= 0 {
			colorFprintf("Monster %v dead, you won !\n", blueString(monster.Name))
			break
		}
	}
	colorFprintf("Combat terminated in %v turn.\n", cyanString(str(turn)))
}

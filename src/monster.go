package main

import (
	"fmt"
	"math/rand"
)

type Monster struct {
	AttackDamage int
	Health       int
	MaxHealth    int
	Name         string
	Race         Race
}

var monster Monster
var isInCombat bool
var isThereMonster = true
var mobNames = []string{"Abrax", "Aeryle", "Afelanidd", "Akvaron", "Almen", "Anjak", "Ariann", "Bassto", "Beele", "Brirelin", "Cambree", "Darusor", "Darzed", "Derle", "Eng", "Ermon", "Flayre", "Gawle", "Glish", "Helsdal", "Hoprig", "Japloon", "Kashtuul", "Kelmerveld", "Kipplob", "Kizarlon", "Leet", "Lizki", "Marb", "Masply", "Merinard", "Mersic", "Milzrik", "Myloryx", "Narvik", "Nesser", "Nyren", "Quelneth", "Quilium", "Quolbin", "Retheer", "Rhiss", "Rience", "Riksul", "Saar", "Saihail", "Shadar", "Simara", "Sounx", "Sraknis", "Syle", "Terragg", "Tourrhok", "Tsai", "Udria", "Vadru", "Varnac", "Varsta", "Viskrek", "Vonir", "Vorshak", "Vryxnir", "Wadziq", "Wryxerg", "Xinsce", "Zamorla", "Zash", "Zeige", "Zheral", "Zoranji"}

func InitGoblin(name string, maxHealth int, attackDamage int) Monster {
	return Monster{
		Name:         name,
		MaxHealth:    maxHealth,
		Health:       maxHealth,
		AttackDamage: attackDamage,
	}
}

func InitMonster(maxHealth int, attackDamage int, race Race) {
	monster = Monster{attackDamage, maxHealth, maxHealth, Random(mobNames), race}
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
	colorPrintf("Monster %v Health\n", redString(monster.showHealth()))
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
		damages = int(float32(boost) / float32(boost/100))
	}
	if boost, ok := monster.Race.Boosts["PoisonResistance"]; weapon.AttackType == Poison && ok {
		damages = int(float32(boost) / float32(boost/100))
	}
	if boost, ok := monster.Race.Boosts["FireResistance"]; weapon.AttackType == Fire && ok {
		damages = int(float32(boost) / float32(boost/100))
	}
	monster.Health -= damages
}

func trainingFight(character *Character, monster *Monster) {
	if !isThereMonster {
		InitMonster(character.getLevel() * 20 + 10 + rand.Intn(30), character.getLevel() + 2 + rand.Intn(3), RandomRace())
		colorPrintf("You search for a new monster...\n%v found, he is a %v with %v max health.", redString(monster.Name), greenString(monster.Race.Name), redString(str(monster.MaxHealth)))
	}

	turn := 0

	isInCombat = true
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
			xp, gold := 5+monster.MaxHealth/10, rand.Intn(20)
			character.gainXP(xp)
			character.Money += gold
			colorPrintf("Monster %v dead, you won %v gold & %v xp!\n", blueString(monster.Name), yellowString(str(gold)), cyanString(str(xp)))
			break
		}
	}
	isInCombat = false
	isThereMonster = false
	colorPrintf("Combat terminated in %v turn.\n", cyanString(str(turn)))
}

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

const trainingMonster = "Training Imperial"

var monster Monster
var isInCombat bool
var isThereMonster = true
var mobNames = []string{"Abrax", "Aeryle", "Afelanidd", "Akvaron", "Almen", "Anjak", "Ariann", "Bassto", "Beele", "Brirelin", "Cambree", "Darusor", "Darzed", "Derle", "Eng", "Ermon", "Flayre", "Gawle", "Glish", "Helsdal", "Hoprig", "Japloon", "Kashtuul", "Kelmerveld", "Kipplob", "Kizarlon", "Leet", "Lizki", "Marb", "Masply", "Merinard", "Mersic", "Milzrik", "Myloryx", "Narvik", "Nesser", "Nyren", "Quelneth", "Quilium", "Quolbin", "Retheer", "Rhiss", "Rience", "Riksul", "Saar", "Saihail", "Shadar", "Simara", "Sounx", "Sraknis", "Syle", "Terragg", "Tourrhok", "Tsai", "Udria", "Vadru", "Varnac", "Varsta", "Viskrek", "Vonir", "Vorshak", "Vryxnir", "Wadziq", "Wryxerg", "Xinsce", "Zamorla", "Zash", "Zeige", "Zheral", "Zoranji"}

func InitMonster(maxHealth int, attackDamage int, race Race, name string) {
	monster = Monster{attackDamage, maxHealth, maxHealth, name, race}
}

func InitRandomMonster(maxHealth int, attackDamage int, race Race) {
	InitMonster(maxHealth, attackDamage, race, Random(mobNames))
}

func (monster *Monster) attack(character *Character) {
	character.Health -= monster.AttackDamage
	printAttack(*monster, *character, monster.AttackDamage)
}

func (monster *Monster) criticalAttack(character *Character) {
	damages := monster.AttackDamage * 2
	character.Health -= damages
	printAttack(*monster, *character, damages)
}

func (monster *Monster) goblinPattern(turn int, character *Character) {
	if turn%3 == 0 {
		monster.criticalAttack(character)
	} else {
		monster.attack(character)
	}
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

func (monster *Monster) printHealth() {
	colorPrintf("Monster %v Health\n", redString(monster.showHealth()))
}

func (monster *Monster) randomPattern(turn int, character *Character) {
	if r := rand.Intn(10); r%3 == 0 {
		monster.criticalAttack(character)
	} else if r%2 == 0 {
		monster.specialAttack(character, RandomAttackType())
	} else {
		monster.attack(character)
	}
}

func (monster *Monster) specialAttack(character *Character, attackType AttackType) {
	item := &Item{
		AttackType: attackType,
	}
	character.HandleAttack(item, monster.AttackDamage)
	printSpecialAttack(*monster, *character, monster.AttackDamage, attackType)
}

func (monster *Monster) showHealth() string {
	return fmt.Sprintf("%v/%v", monster.Health, monster.MaxHealth)
}

func fight(character *Character, monster *Monster) {
	if !isThereMonster {
		InitRandomMonster(character.getLevel()*10+5+rand.Intn(15), character.getLevel()+1+rand.Intn(4), RandomRace())
		colorPrintf("You search for a new monster...\n%v found, he is a %v with %v max health.\n", redString(monster.Name), greenString(monster.Race.Name), redString(str(monster.MaxHealth)))
	}

	if monster.Name == trainingMonster {
		colorPrintf("Hi, I'm Tulius, an Imperial, I'm here to kill you but, it's training, so you can just die, no worries.")
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
		if monster.Name == trainingMonster {
			monster.goblinPattern(turn, character)
		} else {
			monster.randomPattern(turn, character)
		}

		if character.getHealth() <= 0 {
			character.dead()
			break
		}

		if monster.Health <= 0 {
			xp, gold := 5+monster.MaxHealth/10, rand.Intn(20)
			character.gainXP(xp)
			if val, ok := character.RaceBoosts["Money"]; ok {
				gold += gold * (val / 100)
			}
			character.Money += gold
			colorPrintf("Monster %v dead, you won %v gold & %v xp!\n", blueString(monster.Name), yellowString(str(gold)), cyanString(str(xp)))
			break
		}
	}
	isInCombat = false
	isThereMonster = false
	colorPrintf("Combat terminated in %v turn.\n", cyanString(str(turn)))
}

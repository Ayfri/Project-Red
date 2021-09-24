package main

var merchant = Inventory{
	"Health Potions": Item{
		count: 1,
		name:  "Health Potion",
		price: 3,
		onUse: func() {
			character.takeHealthPotion()
		},
	},
	"Poisonous Potion": Item{
		count: 1,
		name:  "Poisonous Potion",
		price: 6,
		onUse: func() {
			character.takePoisonPotion()
		},
	},
	"Spellbook : Fireball": Item{
		count: 1,
		name:  "Spellbook : Fireball",
		price: 25,
		onUse: func() {
			character.spellBook("Spellbook : Fireball")
		},
	},
}

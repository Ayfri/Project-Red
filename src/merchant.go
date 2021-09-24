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
	"Fireball Spellbook": Item{
		count: 1,
		name:  "Fireball Spellbook",
		price: 25,
		onUse: func() {
			character.spellBook("Fireball Spellbook")
		},
	},
	"Wolf Fur": Item{
		count: 1,
		name:  "Wolf Fur",
		price: 4,
	},
	"Troll Skins": Item{
		count: 1,
		name:  "Troll Skin",
		price: 7,
	},
	"Boar's Leathers": Item{
		count: 1,
		name: "Leather's Boar",
		price: 3,
	},
	"Crow's Leathers": Item{
		count: 1,
		name: "Crow's Leather",
		price: 1,
	},
}

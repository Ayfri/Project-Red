package main

var merchant = Inventory{
	"Health Potions": Item{
		count: 1,
		name:  healthPotion,
		price: 3,
		onUse: func(item Item) {
			character.takeHealthPotion()
		},
	},
	"Poisonous Potion": Item{
		count: 1,
		name:  poisonousPotion,
		price: 6,
		onUse: func(item Item) {
			character.takePoisonPotion()
		},
	},
	"Fireball Spellbook": Item{
		count: 1,
		name:  fireballSpellbook,
		price: 25,
		onUse: func(item Item) {
			character.spellBook(fireballSpellbook)
		},
	},
	"Wolf Fur": Item{
		count: 10,
		name:  wolfFur,
		price: 4,
	},
	"Troll Skins": Item{
		count: 10,
		name:  trollSkin,
		price: 7,
	},
	"Boar's Leathers": Item{
		count: 10,
		name: boarFur,
		price: 3,
	},
	"Crow's Leathers": Item{
		count: 10,
		name: crowFeather,
		price: 1,
	},
}

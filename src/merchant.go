package main

var merchant = Inventory{
	"Health Potions": Item{
		Count: 1,
		Name:  healthPotion,
		Price: 3,
		OnUse: func(item Item) {
			character.takeHealthPotion()
		},
	},
	"Poisonous Potion": Item{
		Count: 1,
		Name:  poisonousPotion,
		Price: 6,
		OnUse: func(item Item) {
			character.takePoisonPotion()
		},
	},
	"Fireball Spellbook": Item{
		Count: 1,
		Name:  fireballSpellbook,
		Price: 25,
		OnUse: func(item Item) {
			character.spellBook(fireballSpellbook)
		},
	},
	"Wolf Fur": Item{
		Count: 10,
		Name:  wolfFur,
		Price: 4,
	},
	"Troll Skins": Item{
		Count: 10,
		Name:  trollSkin,
		Price: 7,
	},
	"Boar's Leathers": Item{
		Count: 10,
		Name:  boarFur,
		Price: 3,
	},
	"Crow's Leathers": Item{
		Count: 10,
		Name:  crowFeather,
		Price: 1,
	},
}

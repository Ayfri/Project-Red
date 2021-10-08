package main

type NPC struct {
	CanBuyItems bool
	Inventory   Inventory
	Name        string
	Type        NPCType
}

type NPCType int

const (
	Merchant NPCType = iota
	Blacksmith
)

func (npc *NPC) getInventorySelectorType() SelectorType {
	switch npc.Type {
	case Blacksmith:
		return BlacksmithInventory
	default:
		return MerchantInventory
	}
}

func (npc *NPC) speak() {
	var jobName string
	switch npc.Type {
	case Merchant:
		jobName = "Merchant"
	case Blacksmith:
		jobName = "Blacksmith"
	}
	colorPrintf("Hi stranger, I'm %v, a %v, what do you want ?\n", greenString(npc.Name), yellowString(jobName))
	npc.showMenu()
}

func (npc *NPC) showMenu() {
	var interactionName string
	switch npc.Type {
	case Merchant:
		interactionName = "Buy"
	case Blacksmith:
		 interactionName = "Craft"
	}

	index := 1
	colorPrintf("%v. %v Items\n", cyanString(str(index)), interactionName)

	if npc.CanBuyItems {
		index++
		colorPrintf("%v. Sell Items\n", cyanString(str(index)))
	}
	index++
	colorPrintf("%v. %v\n", cyanString(str(index)), redString("Quit"))

	for {
		input, quit := InputNumber()
		if quit {
			return
		}
		switch input {
		case 1:
			npc.Inventory.makeSelector(npc.getInventorySelectorType(), npc.speak)
		case 2:
			if npc.CanBuyItems {
				character.Inventory.makeSelector(PlayerSellInventory, npc.speak)
			} else {
				return
			}
		case 3:
			return
		}
	}
}

var merchant = NPC{
	CanBuyItems: true,
	Inventory: Inventory{
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
				if isInCombat {
					character.throwPoisonPotion(&monster)
				} else {
					character.takePoisonPotion()
				}
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
	},
	Name: "Lucan",
}

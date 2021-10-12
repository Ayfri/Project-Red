package main

type NPC struct {
	CanBuyItems bool
	Inventory Inventory
	Name      string
	Type      NPCType
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
	var input int
	var quit bool

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

	for !quit && input != 3 {
		input, quit = InputNumber()

		switch input {
		case 1:
			npc.Inventory.makeSelector(npc.getInventorySelectorType(), npc.speak)
			return
		case 2:
			if npc.CanBuyItems {
				character.Inventory.makeSelector(PlayerSellInventory, npc.speak)
				return
			} else {
				return
			}
		}
	}
}

var merchant = NPC{
	CanBuyItems: true,
	Inventory: Inventory{
		healthPotion: Item{
			Count: 1,
			Name:  healthPotion,
			Price: 3,
			OnUse: func(item Item) {
				character.takeHealthPotion()
			},
		},
		poisonousPotion: Item{
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
		inventoryLimit: Item{
			Count: 3,
			Name: inventoryLimit,
			Price: 30,
			OnUse: func(item Item) {
				if character.InventoryLimit == 30 {
					yellow("You can't upgrade your inventory more.")
					return
				}
				character.InventoryLimit += 10
			},
		},
		fireballSpellbook: Item{
			Count: 1,
			Name:  fireballSpellbook,
			Price: 25,
			OnUse: func(item Item) {
				character.spellBook(fireballSpellbook)
			},
		},
		iron: Item{
			Count: 10,
			Name:  iron,
			Price: 4,
		},
		cowLeather: Item{
			Count: 10,
			Name:  cowLeather,
			Price: 7,
		},
		boarFur: Item{
			Count: 10,
			Name:  boarFur,
			Price: 3,
		},
		fiber: Item{
			Count: 10,
			Name:  fiber,
			Price: 1,
		},
	},
	Name: "Lucan",
}

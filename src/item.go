package main

const crowFeather = "Crow Feather"
const boarFur = "Boar Fur"
const trollSkin = "Troll Skin"
const wolfFur = "Wolf Fur"
const healthPotion = "Health Potion"
const poisonousPotion = "Poisonous Potion"
const fireballSpellbook = "Fireball Spellbook"

type Item struct {
	count int
	name  string
	price int
	forgingRequires ForgingRequires
	forgingPrice int
	equipmentType EquipmentType
	onUse func(item Item)
}

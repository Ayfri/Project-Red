package main

const crowFeather = "Crow Feather"
const boarFur = "Boar Fur"
const trollSkin = "Troll Skin"
const wolfFur = "Wolf Fur"
const healthPotion = "Health Potion"
const poisonousPotion = "Poisonous Potion"
const fireballSpellbook = "Fireball Spellbook"

type Item struct {
	Count            int
	Name             string
	Price            int
	ForgingRequires  ForgingRequires
	ForgingPrice     int
	EquipmentType    EquipmentType
	EquipHealthBoost int
	OnUse            func(item Item)
}

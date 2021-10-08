package main

const crowFeather = "Crow Feather"
const boarFur = "Boar Fur"
const trollSkin = "Troll Skin"
const wolfFur = "Wolf Fur"
const healthPotion = "Health Potion"
const poisonousPotion = "Poisonous Potion"
const fireballSpellbook = "Fireball Spellbook"

type Item struct {
	AttackDamage    int
	AttackType       AttackType
	EquipHealthBoost int
	EquipmentType    EquipmentType
	Count            int
	ForgingPrice     int
	ForgingRequires  ForgingRequires
	Name             string
	OnUse            func(item Item)
	Price            int
}

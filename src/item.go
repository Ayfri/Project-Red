package main

const adventurerBoot = "Adventurer's Boots"
const adventurerHat = "Adventurer's Hats"
const adventurerTunic = "Adventurer's Tunics"
const boarFur = "Boar Fur"
const crowFeather = "Crow Feather"
const dragonBane = "Dragonbane"
const fireballSpellbook = "Fireball Spellbook"
const healthPotion = "Health Potion"
const inventoryLimit = "Inventory Limit"
const ironSword = "Iron Sword"
const poisonousPotion = "Poisonous Potion"
const trollSkin = "Troll Skin"
const wolfFur = "Wolf Fur"

type Item struct {
	AttackDamage     int
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

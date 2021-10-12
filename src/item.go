package main

const adventurerBoot = "Adventurer's Boots"
const adventurerHat = "Adventurer's Hats"
const adventurerTunic = "Adventurer's Tunics"
const boarFur = "Boar Fur"
const cowLeather = "Cow Leather"
const dragonBane = "Dragonbane"
const fiber = "Fiber"
const fireballSpellbook = "Fireball Spellbook"
const healthPotion = "Health Potion"
const inventoryLimit = "Inventory Limit"
const ironSword = "Iron Sword"
const iron = "Iron"
const poisonousPotion = "Poisonous Potion"

type Item struct {
	AttackDamage     int
	AttackType       AttackType
	EquipHealthBoost int
	EquipmentType    EquipmentType
	Count            int
	ForgingPrice     int
	ForgingRequires ForgingRequires
	Name            string
	OnUse           func(item Item)
	Price            int
}

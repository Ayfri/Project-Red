package main

var races = []Race{
	{
		Name: "Altmer",
		Boosts: map[string]int{
			"Magic": 50,
		},
	},
	{
		Name:   "Argonian",
		Skills: []string{"WaterBreathing"},
	},
	{
		Name: "Bosmer",
		Boosts: map[string]int{
			"PoisonResistance": 50,
		},
	},
	{
		Name: "Breton",
		Boosts: map[string]int{
			"MagicResistance": 50,
		},
	},
	{
		Name: "Dunmer",
		Boosts: map[string]int{
			"FireResistance": 50,
		},
	},
	{
		Name: "Imperial",
		Boosts: map[string]int{
			"Money": 20,
		},
	},
	{
		Name: "Khajiit",
		Boosts: map[string]int{
			"UnarmedDamages": 15,
		},
	},
	{
		Name: "Nord",
		Boosts: map[string]int{
			"FrostResistance": 50,
		},
	},
	{
		Name:   "Orsimer",
		Skills: []string{"NoAggro"},
	},
	{
		Name: "Redguard",
		Boosts: map[string]int{
			"PoisonResistance": 50,
		},
	},
}

type Race struct {
	Boosts map[string]int
	Name   string
	Skills []string
}

func RaceNames() []string {
	var result []string
	for _, race := range races {
		result = append(result, race.Name)
	}
	return result
}

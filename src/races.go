package main

var races = []Race{
	{
		Name: "Altmer",
		Boosts: map[string]int{
			"MagicResistance": 50,
		},
	},
	{
		Name: "Argonian",
		Boosts : map[string]int {
		    "Money": 15,
		},
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
			"UnarmedDamages": 6,
		},
	},
	{
		Name: "Nord",
		Boosts: map[string]int{
			"FireResistance": 50,
		},
	},
	{
		Name: "Orsimer",
		Boosts: map[string]int{
			"Money": 10,
			"UnarmedDamages": 4,
		},
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
}

func FindRace(name string) (result Race, ok bool) {
	for _, val := range races {
		if val.Name == name {
			return val, true
		}
	}
	return Race{}, false
}

func RaceNames() []string {
	var result []string
	for _, race := range races {
		result = append(result, race.Name)
	}
	return result
}

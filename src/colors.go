package main

import (
	"fmt"
	"github.com/fatih/color"
)

var boldFunc = color.New(color.Bold).PrintlnFunc()
var boldString = color.New(color.Bold).SprintFunc()

func printItemTaken(format string, item string) {
	colorFprintf(format, color.BlueString(item))
}

func printAttack(attacker interface{}, receiver interface{}, damages int) {
	switch receiver.(type) {
	case Monster:
		monster := receiver.(Monster)
		colorFprintf("%v attacked %v, %v damages taken.\n", color.RedString(attacker.(Character).Name), boldString(monster.Name), str(damages))
		monster.printHealth()
	case Character:
		character := receiver.(Character)
		colorFprintf("%v attacked %v, %v damages taken.\n", color.RedString(attacker.(Monster).Name), boldString(character.Name), str(damages))
		character.printHealth()
	}
}

func colorFprintf(format string, vars ...string) {
	result := make([]interface{}, len(vars))
	for i, s := range vars {
		result[i] = s
	}
	fmt.Fprintf(color.Output, format, result...)
}

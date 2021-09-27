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
	colorFprintf("%v attacked %v, %v damages taken.", color.RedString(attacker.(Monster).Name), boldString(receiver.(Monster).Name), str(damages))
	switch receiver.(type) {
	case Monster:
		colorFprintf(receiver.(*Monster).showHealth())
	case Character:
		colorFprintf(receiver.(*Character).showHealth())
	}
}

func colorFprintf(format string, vars ...string) {
	result := make([]interface{}, len(vars))
	for i, s := range vars {
		result[i] = s
	}
	fmt.Fprintf(color.Output, format, result...)
}

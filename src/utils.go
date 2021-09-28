package main

import (
	"fmt"
	"strconv"
	"strings"
)

var lineLength = 60

func Contains(array []interface{}, value interface{}) bool {
	for _, a := range array {
		if a == value {
			return true
		}
	}
	return false
}

func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func (inventory *Inventory) keys() []string {
	keys := make([]string, 0, len(*inventory))
	for k := range *inventory {
		keys = append(keys, k)
	}
	return keys
}

func DifferentKeys(a map[string]int, b Inventory) []string {
	var aKeys []string
	var bKeys []string
	for key := range a {
		aKeys = append(aKeys, key)
	}
	for key := range b {
		bKeys = append(bKeys, key)
	}

	return Difference(aKeys, bKeys)
}

func printLine() {
	fmt.Println(strings.Repeat("-", lineLength))
}

func printCenteredTitle(title string) {
	dashes := strings.Repeat("-", (lineLength-len(title)-2)/2)
	fmt.Printf("%v %v %v\n", dashes, title, dashes)
}

func str(number int) string {
	return strconv.Itoa(number)
}

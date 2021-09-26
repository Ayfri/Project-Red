package main

import "strconv"

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
	for key, _ := range a {
		aKeys = append(aKeys, key)
	}
	for key, _ := range b {
		bKeys = append(bKeys, key)
	}

	return Difference(aKeys, bKeys)
}

func str(number int) string {
	return strconv.Itoa(number)
}
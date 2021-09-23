package main

func Contains(array []interface{}, value interface{}) bool {
	for _, a := range array {
		if a == value {
			return true
		}
	}
	return false
}

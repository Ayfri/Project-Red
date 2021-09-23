package main

import "time"

func poisonPot() {
	for i := 0; i < 3; i++ {
		character.health -= 10
		time.Sleep(time.Second)
	}
}

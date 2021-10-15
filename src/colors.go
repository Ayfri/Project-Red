package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

func InitColors() {
	handle := syscall.Handle(os.Stdout.Fd())
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")
	setConsoleModeProc.Call(uintptr(handle), 0x0001|0x0002|0x0004)
}

func printItemTaken(format string, item string) {
	colorPrintf(format, blueString(item))
}

func printAttack(attacker interface{}, receiver interface{}, damages int) {
	switch receiver.(type) {
	case Monster:
		monster := receiver.(Monster)
		colorPrintf("%v attacked %v, %v damages taken.\n", boldString(attacker.(Character).Name), redString(monster.Name), str(damages))
		monster.printHealth()
	case Character:
		character := receiver.(Character)
		colorPrintf("%v attacked %v, %v damages taken.\n", redString(attacker.(Monster).Name), boldString(character.Name), str(damages))
		character.printHealth()
	}
}

func printSpecialAttack(attacker interface{}, receiver interface{}, damages int, attackType AttackType) {
	var attackName string
	switch attackType {
	case Melee:
		attackName = "melee"
	case Magic:
		attackName = "magic"
	case Poison:
		attackName = "Poison"
	case Fire:
		attackName = "fire"
	}
	switch receiver.(type) {
	case Monster:
		monster := receiver.(Monster)
		colorPrintf("%v attacked %v with a %v attack, %v damages taken.\n", boldString(attacker.(Character).Name), redString(monster.Name), attackName, str(damages))
		monster.printHealth()
	case Character:
		character := receiver.(Character)
		colorPrintf("%v attacked %v with a %v attack, %v damages taken.\n", redString(attacker.(Monster).Name), boldString(character.Name), attackName, str(damages))
		character.printHealth()
	}
}

func colorPrintf(format string, vars ...string) {
	result := make([]interface{}, len(vars))
	for i, s := range vars {
		result[i] = s
	}
	fmt.Printf(format, result...)
}

type Color string

const (
	Reset Color = "\033[0m"

	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	Bold = "\033[1m"
)

func color(color Color, text string) string {
	return fmt.Sprintf("%v%v%v", color, text, Reset)
}

func printColor(c Color, text string, args ...interface{}) {
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	fmt.Printf(color(c, text), args...)
}

func blueString(text string, args ...interface{}) string {
	return color(Blue, fmt.Sprintf(text, args...))
}

func blue(text string, args ...interface{}) {
	if len(args) == 0 {
		printColor(Blue, text)
	} else {
		printColor(Blue, text, args...)
	}
}

func boldString(text string, args ...interface{}) string {
	return color(Bold, fmt.Sprintf(text, args...))
}

func bold(text string, args ...interface{}) {
	if len(args) == 0 {
		printColor(Bold, text)
	} else {
		printColor(Bold, text, args...)
	}
}

func cyanString(text string, args ...interface{}) string {
	return color(Cyan, fmt.Sprintf(text, args...))
}

func cyan(text string, args ...interface{}) {
	if len(args) == 0 {
		printColor(Cyan, text)
	} else {
		printColor(Cyan, text, args...)
	}
}

func greenString(text string, args ...interface{}) string {
	return color(Green, fmt.Sprintf(text, args...))
}

func green(text string, args ...interface{}) {
	if len(args) == 0 {
		printColor(Green, text)
	} else {
		printColor(Green, text, args...)
	}
}

func magentaString(text string, args ...interface{}) string {
	return color(Magenta, fmt.Sprintf(text, args...))
}

func magenta(text string, args ...interface{}) {
	if len(args) == 0 {
		printColor(Magenta, text)
	} else {
		printColor(Magenta, text, args...)
	}
}

func redString(text string, args ...interface{}) string {
	return color(Red, fmt.Sprintf(text, args...))
}

func red(text string, args ...interface{}) {
	if len(args) == 0 {
		printColor(Red, text)
	} else {
		printColor(Red, text, args...)
	}
}

func yellowString(text string, args ...interface{}) string {
	return color(Yellow, fmt.Sprintf(text, args...))
}

func yellow(text string, args ...interface{}) {
	if len(args) == 0 {
		printColor(Yellow, text)
	} else {
		printColor(Yellow, text, args...)
	}
}

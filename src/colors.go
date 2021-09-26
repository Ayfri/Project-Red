package main

import (
	"fmt"
	"github.com/fatih/color"
)

var boldFunc = color.New(color.Bold).PrintlnFunc()
var boldString = color.New(color.Bold).SprintFunc()

func itemTaken(format string, item string) {
	colorFprintf(format, color.BlueString(item))
}

func colorFprintf(format string, vars ...string) {
	result := make([]interface{}, len(vars))
	for i, s := range vars {
		result[i] = s
	}
	fmt.Fprintf(color.Output, format, result...)
}
package main

import (
	"github.com/KlishevichIlya/hexlet-go/greetings"
	"github.com/fatih/color"
)

func main() {
	color.Cyan(greetings.Get())
}

package main

import (
	"fmt"
	form "main/Forms"
	script "main/Script"
	src "main/Sources"
	"github.com/charmbracelet/huh"
)

var (
	Email    string
	Password string
	Start    bool
	Show     bool
	story    string
)

func main() {

	form.FormDisplay(&Email, &Password, &Start, &Show)
	if Start {
		script.ScrriptRunner(Email, Password, src.Steps, Show)
	} else {
		fmt.Println(" Don't Play with me again Right?!")
	}
	JustFun()
}

func JustFun() {

	huh.NewText().
		Title("Are satissfied if you want t change some thing give your feedback!").
		Value(&story).
		Run()
}

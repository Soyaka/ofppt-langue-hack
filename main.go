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
		Title("Are u satissfied? i not keep your feedback!").
		Value(&story).
		Run()
}

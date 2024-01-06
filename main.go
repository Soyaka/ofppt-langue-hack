package main

import (
	"fmt"
	form "main/Forms"
	script "main/Script"
	src "main/Sources"
)

var (
	Email    string
	Password string
	Start    bool
)

func main() {

	form.FormDisplay(&Start, &Email, &Password)
	if Start {
		script.ScrriptRunner(Email, Password, src.Steps)
	}else{
		fmt.Println( " Don't Play with me again Right?!")
}
}

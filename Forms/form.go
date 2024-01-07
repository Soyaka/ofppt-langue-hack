//form.go

package handlers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/huh"
)

func FormDisplay( Email, Password *string, Start ,Show *bool ) {

	form := huh.NewForm(
		huh.NewGroup(
			//email Group
			huh.NewInput().
				Title("enter ur ofppt-langue acount snow man ").
				Value(Email).
				Placeholder("soyaka@ofppt-edu.ma").
				Prompt("👤 ").
				Validate(func(str string) error {
					re := regexp.MustCompile(`@ofppt-edu\.ma`)
					if len(strings.TrimSpace(str)) == 0 {
						return fmt.Errorf("fill it bro")
					}
					if !re.MatchString(str) {
						return fmt.Errorf("is not a valid ofppt e-mail snow woman ?")
					}

					return nil
				}),
		
			// password group
			huh.NewInput().
				Title("now the password  turn focus!").
				Value(Password).
				Prompt("🔒 ").
				Placeholder("focus").
				Password(true),

			huh.NewSelect[bool]().
				Title("would u like to see what happenning ?").
				Options(
					huh.NewOption("yes, show me!",true).Selected(true),
					huh.NewOption("no, keep it abstract",false),
				).
				Value(Show),

			// Ok choice
			huh.NewConfirm().
				Title("would you like to start the the process? ").
				Value(Start),
		),
	)
	form.Run()

}

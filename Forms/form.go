
//form.go

package handlers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/huh"
)

func FormDisplay(Start *bool, Email, Password *string) {

	form := huh.NewForm(
		huh.NewGroup(
			//email Group
			huh.NewInput().
				Title("Enter Ur OFPPT-Langue Acount snow man ! ").
				Value(Email).
				Placeholder("soyaka@ofppt-edu.ma").
				Validate(func(str string) error {
					re := regexp.MustCompile(`@ofppt-edu\.ma`)
					if len(strings.TrimSpace(str)) == 0 {
						return fmt.Errorf("fill it bro")
					}
					if !re.MatchString(str) {
						return fmt.Errorf("is not a valid ofppt e-mail snow woman")
					}

					return nil
				}),
			// password group
			huh.NewInput().
				Title("Now the password Turn focus!").
				Value(Password).
				Placeholder("focus").
				Password(true),

			huh.NewConfirm().
				Title("Would you like to start the the process? ").
				Value(Start),
		),
	)
	form.Run()

}

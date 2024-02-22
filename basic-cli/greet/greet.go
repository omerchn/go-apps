package greet

import (
	"cli/utils"
	"fmt"
)

func RandomGreet(name string) string {
	return fmt.Sprintf(
		utils.GetRandomFromSlice([]string{
			"Hi, %v. Welcome!",
			"Hi there %v.",
			"Hello %v, How's it hanging?",
			"Greetings Mr. %v.",
		}), name,
	)
}

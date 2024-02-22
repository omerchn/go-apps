package main

import (
	"cli/ask"
	"cli/greet"
	"cli/utils"
	"fmt"
)

func main() {
	f_name := ask.Ask("please enter first name: ")
	l_name := ask.Ask("please enter last name: ")

	full_name := utils.ConstructFullName(f_name, l_name)

	fmt.Print(
		"\n",
		"Randomized Greeting:\n",
		greet.RandomGreet(full_name),
		"\n\n",
	)
}

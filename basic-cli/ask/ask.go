package ask

import (
	"fmt"
)

func Ask(prompt string) string {
	fmt.Print(prompt)

	var res string
	fmt.Scanln(&res)

	return res
}

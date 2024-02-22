package lib

import (
	"fmt"
	"os"
)

func GetServerAddr(defPort int) string {
	port, is := os.LookupEnv("PORT")
	if is {
		return fmt.Sprintf(":%v", port)
	} else {
		return fmt.Sprintf(":%v", defPort)
	}
}

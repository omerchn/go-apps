package utils

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetRandomFromSlice[T any](slice []T) T {
	rand.Seed(time.Now().UnixNano())
	return slice[rand.Intn(len(slice))]
}

func ToTitleCase(text string) string {
	return cases.Title(language.English).String(text)
}

func ConstructFullName(first_name string, last_name string) string {
	return fmt.Sprintf("%v %v", ToTitleCase(first_name), ToTitleCase(last_name))
}

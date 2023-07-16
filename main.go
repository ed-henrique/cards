package main

import (
	"fmt"
)

// If the function returns something, it should follow this structure
//
// func FUNCTION_NAME(PARAMETERS) RETURN_TYPE {}
func newCard() string {
	return "Five of Diamonds"
}

// Also, there is no need to import other go files if they are part of the
// same package.

func main() {
	card := newCard()

	fmt.Println(card)
}

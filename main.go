package main

import (
	"fmt"
)

// We can declare variables outside of main function using the following syntax
//
// var NAME TYPE
//
// But we can't use NAME := VALUE
//
// This means that the variable can be initialized without a value before the main function,
// but it can only contain values inside the function itself

func main() {
	// var card string = "Ace of Spades" is the same as below
	//
	// the ":=" infers the type, and initializes the variable with the given value
	card := "Ace of Spades"

	// Remember that, when changing variable values, there's no need to use ":=", since
	// the variable is already initialized
	card = "Five of Diamonds"

	fmt.Println(card)
}

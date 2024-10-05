package main

import (
	"fmt"

	"github.com/ViniciusBatestin/Golang-CLI-Playground/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Vinera")
	fmt.Println(message)
}

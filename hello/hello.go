package main

import (
	"fmt"
	"log"

	"github.com/ViniciusBatestin/Golang-CLI-Playground/greetings"
)

func main() {
	// Set properties of the predefined logger, including
	// the log entre prefix and a flag to disasble printing
	// the time, source file, and line number

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}

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

	// Request a greeting message.
	message, err := greetings.Hello("Vinnys")
	//If an error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console
	fmt.Println(message)
}

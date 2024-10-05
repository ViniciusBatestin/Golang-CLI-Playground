package main

import (
	"fmt"

	"github.com/ViniciusBatestin/Golang-CLI-Playground/tree/master/greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	//Get a greeting message and prin it.
	message := greetings.Hello("Vinnys")
	fmt.Println(message)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the Animal interface
type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

// Struct for Animal characteristics
type Animals struct {
	food       string
	locomotion string
	noise      string
}

// Implement the Animal interface for Animals struct
func (a Animals) Eat() string {
	return a.food
}

func (a Animals) Move() string {
	return a.locomotion
}

func (a Animals) Speak() string {
	return a.noise
}

// Cow, Bird, and Snake struct types embed the Animals struct
type Cow struct {
	Animals
}

type Bird struct {
	Animals
}

type Snake struct {
	Animals
}

// Create a map to store animals by name
var animalsMap = make(map[string]Animal)

// Function to create a new animal based on user input
func newAnimal(name string, animalType string) {
	switch animalType {
	case "cow":
		animalsMap[name] = Cow{Animals{food: "grass", locomotion: "walk", noise: "moo"}}
	case "bird":
		animalsMap[name] = Bird{Animals{food: "worms", locomotion: "fly", noise: "peep"}}
	case "snake":
		animalsMap[name] = Snake{Animals{food: "mice", locomotion: "slither", noise: "hsss"}}
	default:
		fmt.Println("Unknown animal type!")
		return
	}
	fmt.Println("Created it!")
}

// Function to query an animal's action
func queryAnimal(name string, info string) {
	animal, exists := animalsMap[name]
	if !exists {
		fmt.Println("Animal not found!")
		return
	}

	switch info {
	case "eat":
		fmt.Println(animal.Eat())
	case "move":
		fmt.Println(animal.Move())
	case "speak":
		fmt.Println(animal.Speak())
	default:
		fmt.Println("Unknown query! Available options are 'eat', 'move', or 'speak'.")
	}
}

func main() {
	// Set up a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("")
	fmt.Println("Type 'newanimal' [name] [type]. to create a new animal")
	fmt.Println("Type 'query' [name] [info]")

	for {
		// Display prompt
		fmt.Print("> ")

		// Read user input
		scanner.Scan()
		input := scanner.Text()

		// Split the input into parts
		parts := strings.Split(input, " ")

		// Ensure at least 3 parts are provided in the input
		if len(parts) != 3 {
			fmt.Println("Invalid command. Please provide the correct input format.")
			continue
		}

		// Process the 'newanimal' command
		if parts[0] == "newanimal" {
			name := parts[1]
			animalType := parts[2]
			newAnimal(name, animalType)

			// Process the 'query' command
		} else if parts[0] == "query" {
			name := parts[1]
			info := parts[2]
			queryAnimal(name, info)

		} else {
			fmt.Println("Unknown command. Valid commands are 'newanimal' or 'query'.")
		}
	}
}

//When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.

// teh user can create a new animal of one of the three types
// or the user can request informations about an animal that has already created.
// user defines a unique name to the animal
// the user can define animals of a chosen type

// Accepets one command at a time from the user. Print out a response a new prompt on a new line
// every command from the user must be either "newanimal" command or "query" command.

//NEWANIMAL - single line containing 3 str
//1 - newanimal
//2 - name of new animal
//3 - type of the new animal

//"QUERY" - single line contaning 3 str
//1 - "query"
//2 - the name of the animal
// - name of the information requested about the animal "eat" "move" "speak"
// Print out the requested data

//*****
//Define an interface type called Animal which describes the methods of an animal.
//Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
//The Eat() method should print the animal’s food,
//the Move() method should print the animal’s locomotion,
//and the Speak() method should print the animal’s spoken sound.
//Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.

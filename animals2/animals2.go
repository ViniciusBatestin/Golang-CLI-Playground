package main

import (
	"fmt"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Animals struct {
	food       string
	locomotion string
	noise      string
}

// Food eaten methods
func (a Animals) Eat() string {
	return a.food
}

func (a Animals) Move() string {
	return a.locomotion
}

func (a Animals) Speak() string {
	return a.noise
}

// Cow design
type Cow struct{ eat, move, speak string }

func (c Cow) Eat() {
	fmt.Println(c.eat)
}

func (c Cow) Move() {
	fmt.Println(c.move)
}

func (c Cow) Speak() {
	fmt.Println(c.speak)
}

// Bird Design
type Bird struct{ eat, move, speak string }

func (b Bird) Eat() {
	fmt.Println(b.eat)
}

func (b Bird) Move() {
	fmt.Println(b.move)
}

func (b Bird) Speak() {
	fmt.Println(b.speak)
}

// Snake design
type Snake struct{ eat, move, speak string }

func (s Snake) Eat() {
	fmt.Println(s.eat)
}

func (s Snake) Move() {
	fmt.Println(s.move)
}

func (s Snake) Speak() {
	fmt.Println(s.speak)
}

func main() {
	cow := Animals{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animals{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animals{food: "mice", locomotion: "slither", noise: "hsss"}

	fmt.Println("Type a request")
	fmt.Println(">")
	// fmt.Println("reponse goes her")
}

//*****
//Define an interface type called Animal which describes the methods of an animal.
//Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
//The Eat() method should print the animal’s food,
//the Move() method should print the animal’s locomotion,
//and the Speak() method should print the animal’s spoken sound.
//Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.

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

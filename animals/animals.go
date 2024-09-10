package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		var animalName, request string

		fmt.Println("")
		fmt.Println("<Type 'QUIT' to exit>")
		fmt.Println("")
		fmt.Println("************************************************************************************")
		fmt.Println("Type the NAME OF THE ANIMAL 'cow' 'bird' 'snake and ACTION 'move', 'speak' or 'food'")
		fmt.Println("************************************************************************************")
		fmt.Println("")
		fmt.Print(">")
		_, err := fmt.Scanf("%s %s", &animalName, &request)
		if err != nil {
			fmt.Println("Error please enter NAME and ACTION")
		}

		animalName = strings.ToLower(animalName)
		request = strings.ToLower(request)

		if animalName == "quit" {
			fmt.Println("Exiting...")
			break
		}

		var selectedAnimal Animal
		switch animalName {
		case "cow":
			selectedAnimal = cow
		case "bird":
			selectedAnimal = bird
		case "snake":
			selectedAnimal = snake
		default:
			fmt.Println("Unknow Animal")
			return
		}

		switch request {
		case "eat":
			fmt.Println(selectedAnimal.Eat())
		case "move":
			fmt.Println(selectedAnimal.Move())
		case "speak":
			fmt.Println(selectedAnimal.Speak())
		default:
			fmt.Println("Unknow request")
		}
	}
}

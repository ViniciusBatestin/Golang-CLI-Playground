package main

import "fmt"

func Goals(goals ...int) {
	fmt.Println(goals, " ")
	total := 0
	for _, goal := range goals {
		total += goal
	}
	fmt.Println(total)
}

func main() {
	// Goals(40, 30, 7)
	// Goals(2, 3)

	// goals := []int{1, 2, 3}
	// Goals(goals...)
	total := []int{}

	fmt.Println("Enter 3 numbers to sum")
	for i := 0; i < 3; i++ {
		var userInput int
		fmt.Scan(&userInput)
		total = append(total, userInput)
	}

	Goals(total...)

}

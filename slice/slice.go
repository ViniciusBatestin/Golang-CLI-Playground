package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	sorted := make([]int, 0, 3)

	for {
		// handles add input or quit.
		fmt.Println("enters an integer or press X to quit")
		var input string
		fmt.Scan(&input)
		quit := strings.ToUpper(input)
		if quit == "X" {
			break
		}

		//Convert input to int, returns to values
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Not a valid character")
			continue
		}

		//append the input to the slice
		sorted = append(sorted, number)

		//use the Library "sort" to sort the numbers
		sort.Ints(sorted)

		fmt.Println("Your numers sorted are:", sorted)
	}
}

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

Submit your source code for the program,
“slice.go”.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	// Prompt the user for the name of the file
	fmt.Print("What is the name of the file you want to open? ")
	var fileName string
	fmt.Scanln(&fileName)

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a slice to hold the Name structs
	var names []Name

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read each line and split it into first name and last name
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			// Create a Name struct and append it to the slice
			names = append(names, Name{fname: parts[0], lname: parts[1]})
		} else {
			fmt.Println("Skipping line, not formatted correctly:", line)
		}
	}

	// Check for errors encountered by the scanner
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Iterate through the slice and print the first and last names
	fmt.Println("Names found in the file:")
	for _, name := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", name.fname, name.lname)
	}
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Create a map to store user input
	input := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	// Collect user input for name
	fmt.Println("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	// Collect user input for address
	fmt.Println("Enter your address: ")
	scanner.Scan()
	address := scanner.Text()

	// Add the inputs to the map
	input["name"] = name
	input["address"] = address

	// Marshal the map to a JSON object
	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error marshalling JSON", err)
		return
	}

	// Print the JSON object
	fmt.Println("JSON output", string(jsonData))

}

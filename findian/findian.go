package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Type a word so we can make a comparison")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	word := scanner.Text()
	word = strings.ToLower(word)

	found := "Found!"
	notFound := "Not Found!"

	fmt.Println("You type: ", word)

	if strings.Contains(word, "a") && strings.HasPrefix(word, "i") && strings.HasSuffix(word, "n") {
		fmt.Println(found)
	} else {
		fmt.Println(notFound)
	}
}

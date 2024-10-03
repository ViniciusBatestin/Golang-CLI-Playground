package main

import "fmt"

func Mvr(val1, val2 string) (string, string) {
	return val1, val2
}

func main() {
	val1, val2 := Mvr("Mapeka", "Miso")
	fmt.Println(val1, val2)
}

/* 1. Is it possible to return multiple values from a function in Go?
Yes. Multiple values can be returned in Golang by sending comma-separated values with the return statement and by assigning it to multiple variables in a single statement as shown in the example below: */

package main

import (
	"fmt"
	"math"
)

func main(){
	var floatNumber float64
	fmt.Println("Float number")

	fmt.Scan(&floatNumber)

	fmt.Println(math.Trunc(floatNumber))
}
package main

import (
	"fmt"
)

func BubbleSort(sli []int) {
	n := len(sli)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
	// return sli
}

func Swap(sli []int, index int) {
	sli[index], sli[index+1] = sli[index+1], sli[index]
}

func main() {
	var slice []int

	fmt.Println("Please insert up 10 numbers (Or press a Non Integer to get the results)")

	for i := 0; i < 10; i++ {
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Please insert a number")
			break
		}
		slice = append(slice, input)
	}

	fmt.Println("Slice not sorted: ", slice)

	BubbleSort(slice)
	fmt.Println("Sorted slice: ", slice)

}

package main

import (
	"fmt"
	"sort"
	"sync"
)

var numbers []int

func sortSubArray(subArray []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(subArray)
}

func main() {

	fmt.Println("Please input 12 numbers to be sorted:")

	// prompt the user to input a series of int
	for i := 0; i < 12; i++ {
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Please insert a valid number")
			break
		}
		numbers = append(numbers, input)
	}

	var quarterLength = len(numbers) / 4

	// Slicing the slice to partion the slice in 4
	firstQuarter := numbers[0:quarterLength]
	secondQuarter := numbers[quarterLength : 2*quarterLength]
	thirdQuarter := numbers[2*quarterLength : 3*quarterLength]
	fourthQuarter := numbers[3*quarterLength : 4*quarterLength]

	//Working with waitGroup add(), wait() done() avoid race condition
	var wg sync.WaitGroup
	wg.Add(4)

	// 4 different go routines that will printed the subarray sorted, when sorted is completed.
	go sortSubArray(firstQuarter, &wg)
	go sortSubArray(secondQuarter, &wg)
	go sortSubArray(thirdQuarter, &wg)
	go sortSubArray(fourthQuarter, &wg)

	// State waitgroup to wait before and main
	wg.Wait()

	// merge the goroutines sorteed subarrays
	var merginArrays []int
	merginArrays = append(merginArrays, firstQuarter...)
	merginArrays = append(merginArrays, secondQuarter...)
	merginArrays = append(merginArrays, thirdQuarter...)
	merginArrays = append(merginArrays, fourthQuarter...)

	fmt.Println("Fully sorted arrays:", merginArrays)
}

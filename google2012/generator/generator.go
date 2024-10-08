package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string { // return receives-only channel of string.
	c := make(chan string)
	go func() { // we launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // return the channel to the caller
}

func main() {
	c := boring("boring!") // function returning a channel
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
}

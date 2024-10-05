package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // expression to be sent can be any suitable value
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go boring("Boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You day %q\n", <-c) // Receive Expression is just a value.
	}
	fmt.Println("You're boring; I'm leaving")
}

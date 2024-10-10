package main

import (
	"fmt"
	"time"
)

type Chopstick struct {
	id int
}

type Philosopher struct {
	id             int
	leftChopstick  chan *Chopstick
	rightChopstick chan *Chopstick
	eatPermission  chan int
	doneEating     chan int
}

type Host struct {
	permission chan int
}

func (p *Philosopher) eat() {
	for i := 0; i < 3; i++ { // Each philosopher eats 3 times
		// Request permission from the host to eat
		p.eatPermission <- p.id

		// Pick up the chopsticks
		leftChopstick := <-p.leftChopstick
		rightChopstick := <-p.rightChopstick

		// Start eating
		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Second) // Simulate eating time
		fmt.Printf("finishing eating %d\n", p.id)

		// Put down the chopsticks
		p.leftChopstick <- leftChopstick
		p.rightChopstick <- rightChopstick

		// Notify the host that the philosopher is done eating
		p.doneEating <- p.id
	}
}

func (h *Host) managePhilosophers(eatPermission chan int, doneEating chan int) {
	eatingPhilosophers := 0

	for {
		select {
		case philosopherID := <-eatPermission:
			if eatingPhilosophers < 2 {
				eatingPhilosophers++
				fmt.Printf("Host allows philosopher %d to eat\n", philosopherID)
			}

		case philosopherID := <-doneEating:
			eatingPhilosophers--
			fmt.Printf("Host knows philosopher %d finished eating\n", philosopherID)
		}
	}
}

func main() {
	// Create chopstick channels (each chopstick is represented by a channel)
	chopsticks := make([]chan *Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = make(chan *Chopstick, 1)
		chopsticks[i] <- &Chopstick{id: i + 1} // Initialize each chopstick
	}

	// Create host
	host := Host{
		permission: make(chan int),
	}

	// Channels for communication between philosophers and the host
	eatPermission := make(chan int)
	doneEating := make(chan int)

	// Create philosophers
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
			eatPermission:  eatPermission,
			doneEating:     doneEating,
		}
	}

	// Start the host goroutine
	go host.managePhilosophers(eatPermission, doneEating)

	// Start eating routines for each philosopher
	for _, philosopher := range philosophers {
		go philosopher.eat()
	}

	// Prevent the program from exiting too early
	time.Sleep(10 * time.Second)
}

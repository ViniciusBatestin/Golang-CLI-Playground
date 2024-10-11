package main

import (
	"fmt"
	"sync"
	"time"
)

// Philosopher struct representing each philosopher
type Philosopher struct {
	id             int
	leftChopstick  *sync.Mutex
	rightChopstick *sync.Mutex
}

// Host struct to manage eating permissions
type Host struct {
	permission chan struct{} // A channel to limit eating permissions
}

// Function for the philosopher to start eating
func (p *Philosopher) eat(host *Host, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that the philosopher has finished after 3 meals

	for i := 0; i < 3; i++ {
		// Request permission from the host to eat
		host.permission <- struct{}{}

		// Lock the chopsticks (pick them up)
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		// Start eating
		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Second) // Simulate eating time
		fmt.Printf("finishing eating %d\n", p.id)

		// Unlock the chopsticks (put them down)
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		// Notify the host that the philosopher is done eating
		<-host.permission
	}
}

func main() {
	var wg sync.WaitGroup

	// Create the host that manages permissions (allows 2 philosophers to eat at a time)
	host := &Host{
		permission: make(chan struct{}, 2),
	}

	// Create chopsticks (using mutex for each chopstick)
	chopsticks := make([]*sync.Mutex, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = &sync.Mutex{}
	}

	// Create philosophers and assign them chopsticks
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
		}
	}

	// Add philosophers to the wait group (one for each philosopher)
	wg.Add(5)

	// Start eating routines for each philosopher
	for _, philosopher := range philosophers {
		go philosopher.eat(host, &wg)
	}

	// Wait for all philosophers to finish their meals
	wg.Wait()

	fmt.Println("All philosophers have finished eating.")
}

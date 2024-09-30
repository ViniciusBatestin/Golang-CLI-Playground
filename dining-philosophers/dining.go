package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
	host           *Host
}

type Host struct {
	allowedToEat int
	mutex        sync.Mutex
	cond         *sync.Cond
}

func (h *Host) requestPermission() {
	h.mutex.Lock()
	for h.allowedToEat >= 2 {
		h.cond.Wait()
	}
	h.allowedToEat++
	h.mutex.Unlock()
}

func (h *Host) releasePermission() {
	h.mutex.Lock()
	h.allowedToEat--
	h.cond.Signal()
	h.mutex.Unlock()
}

func (p *Philosopher) eat() {
	for i := 0; i < 3; i++ { // Each philosopher eats 3 times
		p.host.requestPermission() // Request permission from the host
		p.leftChopstick.Lock()     // Pick up the left chopstick
		p.rightChopstick.Lock()    // Pick up the right chopstick

		// Start eating
		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Second) // Simulate eating time
		fmt.Printf("finishing eating %d\n", p.id)

		// Release the chopsticks
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		p.host.releasePermission() // Release permission back to the host
	}
}

func main() {
	// Create chopsticks
	chopsticks := make([]*Chopstick, 5)
	for i := range chopsticks {
		chopsticks[i] = &Chopstick{}
	}

	// Create host
	host := &Host{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Create philosophers
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
			host:           host,
		}
	}

	// Start eating routines
	var wg sync.WaitGroup
	for _, philosopher := range philosophers {
		wg.Add(1)
		go func(p *Philosopher) {
			defer wg.Done()
			p.eat()
		}(philosopher)
	}

	// Wait for all philosophers to finish
	wg.Wait()
}

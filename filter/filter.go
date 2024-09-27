// A race condition occurs in concurrent or parallel computing when the behavior of a software system depends on the relative timing or sequence of events,
// such as the order in which threads or processes execute.
// This leads to unpredictable outcomes,
// as the system may produce different results based on the timing of access to shared resources (like memory or variables).
// In simpler terms, a race condition happens when two or more threads/processes attempt to modify a shared resource at the same time,
// and the final outcome depends on the sequence in which these accesses occur.

// In this example we create 2 function that increment that increments a shared var.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Sharred var that both go routines will modify
	var sharedVar int

	// Goroutine 1 increments the shared var
	go func() {
		for i := 0; i < 1000; i++ {
			sharedVar++
			time.Sleep(1 * time.Millisecond)
		}
	}() // Imediatelly invokes the function

	// Goroutine 2: Increments the shared var
	go func() {
		for i := 0; i < 1000; i++ {
			sharedVar++
			time.Sleep(1 * time.Millisecond)
		}
	}()

	time.Sleep(3 * time.Second)

	fmt.Printf("Final value of sharedVar %d\n", sharedVar)
}

// A race condition occurs when two or more goroutines try to modify a shared var concurrently without proper synchorination.
// In this example we create 2 function that increment that increments a shared var.

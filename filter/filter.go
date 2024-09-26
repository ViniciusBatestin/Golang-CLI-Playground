// A race condition occurs when two or more goroutines try to modify a shared var concurrently without proper synchorination.
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

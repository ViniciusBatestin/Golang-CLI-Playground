package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so, t float64
	fmt.Println("Insert acceleration value")
	fmt.Scan(&a)

	fmt.Println("Insert initial velocity value")
	fmt.Scan(&vo)

	fmt.Println("Insert initial displacement")
	fmt.Scan(&so)

	fn := GenDisplaceFn(a, vo, so)

	fmt.Println("Insert value for time")
	fmt.Scan(&t)

	fmt.Printf("The Displacement after %v seconds is %v\n", t, fn(t))

}

/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.
s = ½ a t2 + vot + so
*/

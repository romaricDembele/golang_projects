/*   RACE CONDITION

When we run this program, the 2 goroutines run concurrently and the output is non-deterministic.
Run the program multiple times to see that the output may change.

What is the race condition?: Race condition is a problem where The outcome of the program
depends on the interleavings (non-deterministic ordering).

How it can occur?: It can occur when we use two goroutines.
There is no specific order in which the two goroutines are executed.

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1

	// First goroutine
	go func() {
		x += x
	}()

	// Second go routine
	go func() {
		fmt.Println(x)
	}()

	// Wait a second before exiting the program
	time.Sleep(time.Second)
}

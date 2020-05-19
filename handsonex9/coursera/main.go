//race conditons occur due to communication .

package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	c := 0

	for i := 0; i < 2; i++ {
		go func() { // here the function establishes a goroutine.
					// Thus whatever is in go func() block will execute when the machine communicates.

			fmt.Println("func", i)
			c++                   // the counter value could be between 0 to 5 times due to race conditons.
								  // Race conditon occurs as the func block gets off to a goroutine .
								  // hence the statements occur in a non - deterministic order.

		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter", c)
}

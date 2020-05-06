package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	wg.Add(1)

	go func() {

		c <- 42
		wg.Wait()
	}()

	wg.Done()
	fmt.Println(<-c)
}

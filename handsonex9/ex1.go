package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Add(2)

	go foo()

	go bar()

	wg.Wait()

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()

}

func bar() {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
	wg.Done()
}

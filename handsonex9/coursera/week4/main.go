package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type chops struct {
	sync.Mutex
}

type philos struct {
	number          int
	leftcs, rightcs *chops
}

func (p philos) eat() {

	for i := 0; i < 3; i++ {
		p.leftcs.Lock()
		p.rightcs.Lock()

		fmt.Println("Starting to eat Philosopher", p.number)

		p.leftcs.Unlock()
		p.rightcs.Unlock()
		fmt.Println("Finished eating Philosopher", p.number)
		time.Sleep(time.Second)
	}
	wg.Done()

}
func main() {

	Cs := make([]*chops, 5, 5)
	for i := 0; i < 5; i++ {
		Cs[i] = new(chops)
	}
	philo := make([]*philos, 5, 5)
	for j := 0; j < 5; j++ {
		philo[j] = &philos{j, Cs[j], Cs[(j+1)%5]}
	}
	for k := 0; k < 5; k++ {
		wg.Add(1)
		go philo[k].eat()
	}

	wg.Wait()
}

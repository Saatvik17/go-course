//four goroutines used to sort the numbers.
// slice used with array as underlying type.

package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg sync.WaitGroup

func main() {

	var len int
	fmt.Println("Enter the length of array")
	fmt.Scan(&len)
	sli := make([]int, len, len)

	for i := range sli {
		fmt.Println("Enter element", i)
		fmt.Scan(&sli[i])

	}
	wg.Add(4)
	sli1 := sli[0:(len / 4)]
	sli2 := sli[(len / 4):(len / 2)]
	sli3 := sli[(len / 2):(3 * (len / 4))]
	sli4 := sli[(3 * (len / 4)):len]

	go group1(sli1)
	go group2(sli2)
	go group3(sli3)
	go group4(sli4)

	slinew := make([]int, 0, len)

	slinew = append(slinew, sli1...)
	slinew = append(slinew, sli2...)
	slinew = append(slinew, sli3...)
	slinew = append(slinew, sli4...)

	wg.Wait()

	sort.Ints(slinew)
	fmt.Println("sorted and combined ", slinew)

}

func group1(xi []int) {
	sort.Ints(xi)
	fmt.Println("sorting group 1 ", xi)
	wg.Done()

}

func group2(xi []int) {
	sort.Ints(xi)
	fmt.Println("sorting group 2 ", xi)
	wg.Done()
}

func group3(xi []int) {
	sort.Ints(xi)
	fmt.Println("sorting group 3 ", xi)
	wg.Done()
}
func group4(xi []int) {
	sort.Ints(xi)
	fmt.Println("sorting group 4 ", xi)
	wg.Done()
}

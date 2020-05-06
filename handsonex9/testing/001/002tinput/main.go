package main

import "fmt"

func main() {

	var input1 int
	var input2 int
	fmt.Println("enter your numbers")
	fmt.Scan(&input1)
	fmt.Scan(&input2)

	fmt.Println("sum is ", add(input1, input2))

}

func add(x, y int) int {

	s := x + y
	return s
}

package main

import "fmt"

func main() {

	fmt.Println("difference is ", diff(2, 3))

}

func diff(x, y int) int {

	t := y - x

	return t
}

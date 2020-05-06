package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	var input1 int
	var input2 int
	fmt.Scan(&input1)
	fmt.Scan(&input2)
	y := add(input1, input2)
	if y != add(input1, input2) {
		fmt.Println("error")
	}
}

func BenchmarkAdd(b *testing.B) {
	var input1 int
	var input2 int
	fmt.Scan(&input1)
	fmt.Scan(&input2)
	for i := 0; i <= b.N; i++ {
		add(input1, input2)
	}
}

func Exampleadd() {

	fmt.Println(add(2, 3))
	//Output:
	//5
}

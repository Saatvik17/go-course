package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {

	p1 := person{" james", "bond"}

	saysomething(&p1)

}

func (p1 *person) speak() {
	fmt.Println("hola")
}

type human interface {
	speak()
}

func saysomething(h human) {

	h.speak()
}

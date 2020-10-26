package main

import (
	"fmt"
)

func addition(a int, b int) int {
	fmt.Println("addition has called")
	return a + b
}

func compute() int {
	n := &Numbers{1, 2}
	interfaceAdd := n.addition()

	pureAdd := addition(1, 1)

	fmt.Println("interfaceAdd", interfaceAdd)
	fmt.Println("pureAdd", pureAdd)

	return interfaceAdd - 2
}

// Numbers struct
type Numbers struct {
	A int
	B int
}

type operations interface {
	addition()
}

func (n Numbers) addition() int {
	fmt.Println("n.addition has called")

	return 20
}

func (n Numbers) compute() int {
	fmt.Println("n.compute has called")

	return 50
}

func main() {
	nb := compute()
	fmt.Println(nb)
}

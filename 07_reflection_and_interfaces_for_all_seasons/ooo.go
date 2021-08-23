package main

import "fmt"

type aa struct {
	XX int
	YY int
}

func (A aa) A() {
	fmt.Println("Function A() for A")
}

type bb struct {
	AA string
	XX int
}

func (B bb) A() {
	fmt.Println("Function A() for B")
}

type cc struct {
	A aa
	B bb
}

func main() {
	var i cc
	i.A.A()
	i.B.A()
}

package main

import (
	"fmt"
	"math"

	"mastering-go/07_reflection_and_interfaces_for_all_seasons/myInterface"
)

type square2 struct {
	X float64
}

type circle2 struct {
	R float64
}

type rectangle struct {
	X float64
	Y float64
}

func (s square2) Area() float64 {
	return s.X * s.X
}

func (s square2) Perimeter() float64 {
	return 4 * s.X
}

func (s circle2) Area() float64 {
	return s.R * s.R * math.Pi
}

func (s circle2) Perimeter() float64 {
	return 2 * s.R * math.Pi
}

func tellInterface(x interface{}) {
	switch v := x.(type) {
	case square2:
		fmt.Println("This is a square!")
	case circle2:
		fmt.Printf("%v is a circle!\n", v)
	case rectangle:
		fmt.Println("This is a rectangle!")
	default:
		fmt.Printf("Unknown type %T!\n", v)
	}
}

func Calculate2(x myInterface.Shape) {
	_, ok := x.(circle2)
	if ok {
		fmt.Println("Is a circle2!")
	}

	v, ok := x.(square2)
	if ok {
		fmt.Println("Is a square2:", v)
	}

	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

func main() {
	x := circle2{R: 10}
	tellInterface(x)
	y := rectangle{X: 4, Y: 1}
	tellInterface(y)
	z := square2{X: 4}
	tellInterface(z)
	tellInterface(10)
}

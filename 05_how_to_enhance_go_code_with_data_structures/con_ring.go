package main

import (
	"container/ring"
	"fmt"
)

var sizeOfRing int = 10

func main() {
	myRing := ring.New(sizeOfRing + 1)
	fmt.Println("Empty ring:", myRing)

	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}

	myRing.Value = 2

	sum := 0
	myRing.Do(func(x interface{}) {
		t := x.(int)
		sum = sum + t
	})
	fmt.Println("Sum:", sum)
}

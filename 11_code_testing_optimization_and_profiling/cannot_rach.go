package main

import "fmt"

func f11() int {
	fmt.Println("Entering f1()")
	return -10
	fmt.Println("Exiting f1()")
	return -1
}
func f22() int {
	if true {
		return 10
	}
	fmt.Println("Exiting f2()")
	return 0
}
func main() {
	fmt.Println(f11())
	fmt.Println("Exiting program...")
}

package main

import "C"
import "fmt"

//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

// go build -o used_by_c.o -buildmode=c-shared used_by_c.go
func main() {}

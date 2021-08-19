package main

import (
	"fmt"

	"mastering-go/06_what_you_might_not_know_about_go_packages_and_functions/a"
	"mastering-go/06_what_you_might_not_know_about_go_packages_and_functions/b"
)

func init() {
	fmt.Println("init() manyInit")
}

func main() {
	a.FromA()
	b.FromB()
}

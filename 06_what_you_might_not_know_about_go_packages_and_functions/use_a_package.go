package main

import (
	"fmt"

	"mastering-go/06_what_you_might_not_know_about_go_packages_and_functions/aPackage"
)

func main() {
	fmt.Println("Using aPackage!")
	aPackage.A()
	aPackage.B()
	fmt.Println(aPackage.MyConstant)
}

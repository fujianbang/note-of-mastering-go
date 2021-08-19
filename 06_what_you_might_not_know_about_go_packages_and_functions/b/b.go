package b

import (
	"fmt"

	"mastering-go/06_what_you_might_not_know_about_go_packages_and_functions/a"
)

func init() {
	fmt.Println("init() b")
}

func FromB() {
	fmt.Println("fromB()")
	a.FromA()
}

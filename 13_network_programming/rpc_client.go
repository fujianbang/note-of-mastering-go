package main

import (
	"fmt"
	"net/rpc"
	"os"

	"mastering-go/13_network_programming/sharedRPC"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string!")
		return
	}

	connect := arguments[1]
	c, err := rpc.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := sharedRPC.MyFloats{A1: 16, A2: -0.5}
	var reply float64
	if err := c.Call("MyInterface.Multiply", args, &reply); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %f\n", reply)

	if err := c.Call("MyInterface.Power", args, &reply); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)
}

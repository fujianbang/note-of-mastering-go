package main

import (
	"fmt"
	"math"
	"net"
	"net/rpc"
	"os"

	"mastering-go/13_network_programming/sharedRPC"
)

type MyInterface struct{}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func (t *MyInterface) Multiply(arguments *sharedRPC.MyFloats, reply *float64) error {
	*reply = arguments.A1 * arguments.A2
	return nil
}

func (t *MyInterface) Power(arguments *sharedRPC.MyFloats, reply *float64) error {
	*reply = Power(arguments.A1, arguments.A2)
	fmt.Println(*reply)
	return nil
}

func main() {
	port := ":1234"
	arguments := os.Args
	if len(arguments) != 1 {
		port = ":" + arguments[1]
	}

	myInterface := new(MyInterface)
	rpc.Register(myInterface)
	t, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}

}

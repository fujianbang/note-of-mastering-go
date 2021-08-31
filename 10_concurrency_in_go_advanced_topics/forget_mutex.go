package main

import (
	"fmt"
	"sync"
)

var m1 sync.Mutex

func function() {
	// DON'T FORGET UNLOCK() MUTEX
	// defer m1.Unlock()
	m1.Lock()
	fmt.Println("Locked!")
}

func main() {
	var w sync.WaitGroup

	go func() {
		defer w.Done()
		function()
	}()
	w.Add(1)

	go func() {
		defer w.Done()
		function()
	}()
	w.Add(1)

	w.Wait()
}

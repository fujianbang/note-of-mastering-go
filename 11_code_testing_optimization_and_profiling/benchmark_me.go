package main

import "fmt"

func fibo111(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo111(n-1) + fibo111(n-2)
	}
}

func fibo222(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibo222(n-1) + fibo222(n-2)
}

func fibo333(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	fmt.Println(fibo111(40))
	fmt.Println(fibo222(40))
	fmt.Println(fibo333(40))
}

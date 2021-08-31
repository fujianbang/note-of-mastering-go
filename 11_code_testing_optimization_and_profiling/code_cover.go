package main

func fibo11(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo11(n-1) + fibo11(n-2)
	}
}
func fibo22(n int) int {
	if n >= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo11(n-1) + fibo11(n-2)
	}
}

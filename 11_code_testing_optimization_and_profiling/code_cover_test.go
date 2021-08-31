package main

import "testing"

func TestFibo1(t *testing.T) {
	if fibo11(1) != 1 {
		t.Errorf("Error fibo1(1): %d\n", fibo11(1))
	}
}
func TestFibo2(t *testing.T) {
	if fibo22(0) != 0 {
		t.Errorf("Error fibo2(0): %d\n", fibo11(0))
	}
}
func TestFibo1_10(t *testing.T) {
	if fibo11(10) == 1 {
		t.Errorf("Error fibo1(1): %d\n", fibo11(1))
	}
}
func TestFibo2_10(t *testing.T) {
	if fibo22(10) != 0 {
		t.Errorf("Error fibo2(0): %d\n", fibo11(0))
	}
}

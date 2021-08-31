package main

import (
	"testing"
	"testing/quick"
)

var N = 1000000

func TestWithSystem(t *testing.T) {
	condition := func(a, b uint16) bool {
		return Add(a, b) == (b + a)
	}
	if err := quick.Check(condition, &quick.Config{MaxCount: N}); err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestWithItself(t *testing.T) {
	condition := func(a, b uint16) bool {
		return Add(a, b) == Add(b, a)
	}

	if err := quick.Check(condition, &quick.Config{MaxCount: N}); err != nil {
		t.Errorf("Error: %v", err)
	}
}

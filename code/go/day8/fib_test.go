package fib

import (
	"testing"
)

func TestFib(t *testing.T) {
	if fib(3) != 2 {
		t.Errorf("want 2")
	}
	if fib(5) != 5 {
		t.Errorf("want 5")
	}
	if fib(6) != 8 {
		t.Errorf("want 8")
	}
	if fib(7) != 13 {
		t.Errorf("want 13")
	}
}
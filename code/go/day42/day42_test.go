package day42

import "testing"

func TestFunc(t *testing.T) {
	if NumberOf1Between1AndN(1) != 1 {
		t.Errorf("not correct.")
	}

	if NumberOf1Between1AndN(5) != 1 {
		t.Errorf("not correct.")
	}

	if NumberOf1Between1AndN(10) != 2 {
		t.Errorf("not correct.")
	}

	if NumberOf1Between1AndN(55) != 16 {
		t.Errorf("not correct.")
	}

	if NumberOf1Between1AndN(99) != 20 {
		t.Errorf("not correct.")
	}

	if NumberOf1Between1AndN(10000) != 4001 {
		t.Errorf("not correct.")
	}

	if NumberOf1Between1AndN(0) != 0 {
		t.Errorf("not correct.")
	}
}

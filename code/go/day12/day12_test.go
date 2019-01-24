package day12

import "testing"

func TestSolution1(t *testing.T) {
	if NumberOf1Solution1(0) != 0 {
		t.Errorf("not pass.")
	}

	if NumberOf1Solution1(1) != 1 {
		t.Errorf("not pass.")
	}

	if NumberOf1Solution1(10) != 2 {
		t.Errorf("not pass.")
	}

	if NumberOf1Solution1(0x7FFFFFFF) != 31 {
		t.Errorf("not pass.")
	}

	if NumberOf1Solution1(0xFFFFFFFF) != 32 {
		t.Errorf("not pass.")
	}

	if NumberOf1Solution1(0x80000000) != 1 {
		t.Errorf("not pass.")
	}
}

func TestSolution2(t *testing.T) {
	if NumberOf1Solution2(0) != 0 {
		t.Errorf("not pass.")
	}

	if NumberOf1Solution2(1) != 1 {
		t.Errorf("not pass.")
	}

	if NumberOf1Solution2(10) != 2 {
		t.Errorf("not pass.")
	}

	if NumberOf1Solution2(0x7FFFFFFF) != 31 {
		t.Errorf("not pass.")
	}

	if NumberOf1Solution2(0xFFFFFFFF) != 32 {
		t.Errorf("not pass.")
	}

	if NumberOf1Solution2(0x80000000) != 1 {
		t.Errorf("not pass.")
	}
}
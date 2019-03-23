package day43

import "testing"

func TestFunc(t *testing.T) {
	if DigitAtIndex(0) != 0 {
		t.Errorf("not correct.")
	}

	if DigitAtIndex(1) != 1 {
		t.Errorf("not correct.")
	}

	if DigitAtIndex(9) != 9 {
		t.Errorf("not correct.")
	}

	if DigitAtIndex(10) != 1 {
		t.Errorf("not correct.")
	}

	if DigitAtIndex(189) != 9 {
		t.Errorf("not correct.")
	}

	if DigitAtIndex(190) != 1 {
		t.Errorf("not correct.")
	}

	if DigitAtIndex(1000) != 3 {
		t.Errorf("not correct.")
	}

	if DigitAtIndex(1001) != 7 {
		t.Errorf("not correct.")
	}

	if DigitAtIndex(1002) != 0 {
		t.Errorf("not correct.")
	}
}
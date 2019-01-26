package day14

import "testing"

func TestPow(t *testing.T) {
	if Pow(2, 3) != 8 {
		t.Errorf("not correct.")
	}

	if Pow(-2, 3) != -8 {
		t.Errorf("not correct.")
	}

	if Pow(2, -3) != 0.125 {
		t.Errorf("not correct.")
	}

	if Pow(2, 0) != 1 {
		t.Errorf("not correct.")
	}

	if Pow(0, 0) != 0 {
		t.Errorf("not correct.")
	}

	if Pow(0, 4) != 0 {
		t.Errorf("not correct.")
	}

	if Pow(0, -4) != 0 {
		t.Errorf("not correct.")
	}
}

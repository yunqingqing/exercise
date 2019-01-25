package day13

import "testing"

func TestMaxProductAfterCutting1(t *testing.T) {
	if MaxProductAfterCyttingSolution1(1) != 0 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution1(2) != 1 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution1(3) != 2 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution1(4) != 4 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution1(5) != 6 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution1(6) != 9 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution1(7) != 12 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution1(8) != 18 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution1(9) != 27 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution1(10) != 36 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution1(50) != 86093442 {
		t.Errorf("not correct.")
	}
}

func TestMaxProductAfterCutting2(t *testing.T) {
	if MaxProductAfterCyttingSolution2(1) != 0 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution2(2) != 1 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution2(3) != 2 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution2(4) != 4 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution2(5) != 6 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution2(6) != 9 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution2(7) != 12 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution2(8) != 18 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution2(9) != 27 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution2(10) != 36 {
		t.Errorf("not correct.")
	}

	if MaxProductAfterCyttingSolution2(50) != 86093442 {
		t.Errorf("not correct.")
	}
}
package day38

import "testing"

func TestFunc1(t *testing.T) {
	array := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}

	if MoreThanHalfNum(array, len(array)) != 2 {
		t.Errorf("not correct.")
	}
}

func TestFunc2(t *testing.T) {
	array := []int{1, 2, 3, 2, 4, 2, 5, 2, 3}

	if MoreThanHalfNum(array, len(array)) != 0 {
		t.Errorf("not correct.")
	}
}

func TestFunc3(t *testing.T) {
	array := []int{2, 2, 2, 2, 2, 1, 3, 4, 5}

	if MoreThanHalfNum(array, len(array)) != 2 {
		t.Errorf("not correct.")
	}
}

func TestFunc4(t *testing.T) {
	array := []int{1, 3, 4, 5, 2, 2, 2, 2, 2}

	if MoreThanHalfNum(array, len(array)) != 2 {
		t.Errorf("not correct.")
	}
}

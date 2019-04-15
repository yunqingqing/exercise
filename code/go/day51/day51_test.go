package day51

import "testing"

func TestFunc1(t *testing.T) {
	data := []int{1, 2, 3, 3, 3, 3, 4, 5}

	if GetNumberOfK(data, len(data), 3) != 4 {
		t.Errorf("not correct.")
	}
}

func TestFunc2(t *testing.T) {
	data := []int{3, 3, 3, 3, 4, 5}

	if GetNumberOfK(data, len(data), 3) != 4 {
		t.Errorf("not correct.")
	}
}

func TestFunc3(t *testing.T) {
	data := []int{1, 2, 3, 3, 3, 3}

	if GetNumberOfK(data, len(data), 3) != 4 {
		t.Errorf("not correct.")
	}
}

func TestFunc4(t *testing.T) {
	data := []int{1, 3, 3, 3, 3, 4, 5}

	if GetNumberOfK(data, len(data), 2) != 0 {
		t.Errorf("not correct.")
	}
}

func TestFunc5(t *testing.T) {
	data := []int{1, 3, 3, 3, 3, 4, 5}

	if GetNumberOfK(data, len(data), 0) != 0 {
		t.Errorf("not correct.")
	}
}

func TestFunc6(t *testing.T) {
	data := []int{1, 3, 3, 3, 3, 4, 5}

	if GetNumberOfK(data, len(data), 6) != 0 {
		t.Errorf("not correct.")
	}
}

func TestFunc7(t *testing.T) {
	data := []int{3, 3, 3, 3}

	if GetNumberOfK(data, len(data), 3) != 4 {
		t.Errorf("not correct.")
	}
}

func TestFunc8(t *testing.T) {
	data := []int{3, 3, 3, 3}

	if GetNumberOfK(data, len(data), 4) != 0 {
		t.Errorf("not correct.")
	}
}

func TestFunc9(t *testing.T) {
	data := []int{3}

	if GetNumberOfK(data, len(data), 3) != 1 {
		t.Errorf("not correct.")
	}
}
package day49

import "testing"

func TestFunc1(t *testing.T) {
	data := []int{1, 2, 3, 4, 7, 6, 5}
	expected := 3
	if InversePairs(data, len(data)) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc2(t *testing.T) {
	data := []int{6, 5, 4, 3, 2, 1}
	expected := 15
	if InversePairs(data, len(data)) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc3(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	expected := 0
	if InversePairs(data, len(data)) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc4(t *testing.T) {
	data := []int{1}
	expected := 0
	if InversePairs(data, len(data)) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc5(t *testing.T) {
	data := []int{1, 2}
	expected := 0
	if InversePairs(data, len(data)) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc6(t *testing.T) {
	data := []int{2, 1}
	expected := 1
	if InversePairs(data, len(data)) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc7(t *testing.T) {
	data := []int{1, 2, 1, 2, 1}
	expected := 3
	if InversePairs(data, len(data)) != expected {
		t.Errorf("not correct.")
	}
}

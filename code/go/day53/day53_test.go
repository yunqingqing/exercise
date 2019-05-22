package day53

import "testing"

func TestFunc1(t *testing.T) {
	var num1, num2 int
	data := []int{2, 4, 3, 6, 3, 2, 5, 5}
	FindNumsAppearOnce(data, len(data), &num1, &num2)
	expected1 := 4
	expected2 := 6
	if (num1 != expected1 && num2 != expected2) || (num1 != expected2 && num2 != expected1) {
		t.Errorf("not correct.")
	}
}

func TestFunc2(t *testing.T) {
	var num1, num2 int
	data := []int{4, 6}
	FindNumsAppearOnce(data, len(data), &num1, &num2)
	expected1 := 4
	expected2 := 6
	if (num1 != expected1 && num2 != expected2) || (num1 != expected2 && num2 != expected1) {
		t.Errorf("not correct.")
	}
}

func TestFunc3(t *testing.T) {
	var num1, num2 int
	data := []int{4, 6, 1, 1, 1, 1}
	FindNumsAppearOnce(data, len(data), &num1, &num2)
	expected1 := 4
	expected2 := 6
	if (num1 != expected1 && num2 != expected2) || (num1 != expected2 && num2 != expected1) {
		t.Errorf("not correct.")
	}
}
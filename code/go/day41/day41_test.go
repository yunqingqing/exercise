package day41

import "testing"

func TestFunc1(t *testing.T) {
	nums := []int{1, -2, 3, 10, -4, 7, 2, -5}
	if FindGreatestSumOfSubArray(nums, len(nums)) != 18 {
		t.Errorf("not correct.")
	}
}

func TestFunc2(t *testing.T) {
	nums := []int{-2, -8, -1, -5, -9}
	if FindGreatestSumOfSubArray(nums, len(nums)) != -1 {
		t.Errorf("not correct.")
	}
}

func TestFunc3(t *testing.T) {
	nums := []int{2, 8, 1, 5, 9}
	if FindGreatestSumOfSubArray(nums, len(nums)) != 25 {
		t.Errorf("not correct.")
	}
}

func TestFunc4(t *testing.T) {
	nums := []int{}
	if FindGreatestSumOfSubArray(nums, len(nums)) != 0 {
		t.Errorf("not correct.")
	}
}

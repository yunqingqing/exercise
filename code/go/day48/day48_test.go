package day48

import "testing"

func TestFunc1(t *testing.T) {
	str := "google"
	if FirstNotRepeatingChar(str) != "l" {
		t.Errorf("not correct.")
	}
}

func TestFunc2(t *testing.T) {
	str := "aabccdbd"
	if FirstNotRepeatingChar(str) != "" {
		t.Errorf("not correct.")
	}
}

func TestFunc3(t *testing.T) {
	str := "abcdefg"
	if FirstNotRepeatingChar(str) != "a" {
		t.Errorf("not correct.")
	}
}

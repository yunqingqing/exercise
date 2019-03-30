package day47

import "testing"

func TestFunc1(t *testing.T) {
	str := "abcacfrar"
	expected := 4
	if LongestSubstringWithoutDuplication(str) != expected {
		t.Errorf("not correct")
	}
}

func TestFunc2(t *testing.T) {
	str := "acfrarabc"
	expected := 4
	if LongestSubstringWithoutDuplication(str) != expected {
		t.Errorf("not correct")
	}
}

func TestFunc3(t *testing.T) {
	str := "arabcacfr"
	expected := 4
	if LongestSubstringWithoutDuplication(str) != expected {
		t.Errorf("not correct")
	}
}

func TestFunc4(t *testing.T) {
	str := "aaaa"
	expected := 1
	if LongestSubstringWithoutDuplication(str) != expected {
		t.Errorf("not correct")
	}
}

func TestFunc5(t *testing.T) {
	str := "abcdefg"
	expected := 7
	if LongestSubstringWithoutDuplication(str) != expected {
		t.Errorf("not correct")
	}
}

func TestFunc6(t *testing.T) {
	str := "aaabbbccc"
	expected := 2
	if LongestSubstringWithoutDuplication(str) != expected {
		t.Errorf("not correct")
	}
}

func TestFunc7(t *testing.T) {
	str := "abcdcba"
	expected := 4
	if LongestSubstringWithoutDuplication(str) != expected {
		t.Errorf("not correct")
	}
}

func TestFunc8(t *testing.T) {
	str := "abcdaef"
	expected := 6
	if LongestSubstringWithoutDuplication(str) != expected {
		t.Errorf("not correct")
	}
}

func TestFunc9(t *testing.T) {
	str := "a"
	expected := 1
	if LongestSubstringWithoutDuplication(str) != expected {
		t.Errorf("not correct")
	}
}

func TestFunc10(t *testing.T) {
	str := ""
	expected := 0
	if LongestSubstringWithoutDuplication(str) != expected {
		t.Errorf("not correct")
	}
}

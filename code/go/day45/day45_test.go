package day45

import "testing"

func TestFunc1(t *testing.T) {
	number := 10
	expected := 2
	if GetTranslationCount(number) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc2(t *testing.T) {
	number := 0
	expected := 1
	if GetTranslationCount(number) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc3(t *testing.T) {
	number := 126
	expected := 2
	if GetTranslationCount(number) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc4(t *testing.T) {
	number := 426
	expected := 1
	if GetTranslationCount(number) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc5(t *testing.T) {
	number := 100
	expected := 2
	if GetTranslationCount(number) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc6(t *testing.T) {
	number := 101
	expected := 2
	if GetTranslationCount(number) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc7(t *testing.T) {
	number := 12258
	expected := 5
	if GetTranslationCount(number) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc8(t *testing.T) {
	number := -100
	expected := 0
	if GetTranslationCount(number) != expected {
		t.Errorf("not correct.")
	}
}

func TestFunc9(t *testing.T) {
	number := 125
	expected := 3
	if GetTranslationCount(number) != expected {
		t.Errorf("not correct.")
	}
}

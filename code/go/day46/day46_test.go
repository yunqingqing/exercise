package day46

import "testing"

func TestFunc1(t *testing.T) {
	values := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	excepted := 29
	
	if GetMaxValue(values, 3, 3) != excepted {
		t.Errorf("not correct.")
	}
}

func TestFunc2(t *testing.T) {
	values := [][]int{
        { 1, 10, 3, 8 },
        { 12, 2, 9, 6 },
        { 5, 7, 4, 11 },
        { 3, 7, 16, 5 },
	}
	excepted := 53
	
	if GetMaxValue(values, 4, 4) != excepted {
		t.Errorf("not correct.")
	}
}

func TestFunc3(t *testing.T) {
	values := [][]int{
		{ 1, 10, 3, 8 },
	}
	excepted := 22
	
	if GetMaxValue(values, 1, 4) != excepted {
		t.Errorf("not correct.")
	}
}

func TestFunc4(t *testing.T) {
	values := [][]int{
        { 1 },
        { 12 },
        { 5 },
        { 3 },
	}
	excepted := 21
	
	if GetMaxValue(values, 4, 1) != excepted {
		t.Errorf("not correct.")
	}
}

func TestFunc5(t *testing.T) {
	values := [][]int{
        { 3 },
	}
	excepted := 3
	
	if GetMaxValue(values, 1, 1) != excepted {
		t.Errorf("not correct.")
	}
}
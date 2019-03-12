package day40

import (
	"testing"
)

func TestFunc(t *testing.T) {
	Insert(2)
	Insert(4)
	Insert(6)
	Insert(8)
	Insert(10)
	if GetMedian() != 6 {
		t.Errorf("not correct.")
	}
	Insert(12)
	if GetMedian() != 7 {
		t.Errorf("not correct.")
	}
}

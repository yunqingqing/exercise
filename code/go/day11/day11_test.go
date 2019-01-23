package day11

import (
	"testing"
)

func TestFunc(t *testing.T) {
	if movingCount(5, 10, 10) != 21 {
		t.Errorf("count err.")
	}

	if movingCount(15, 20, 20) != 359 {
		t.Errorf("count err.")
	}

	if movingCount(10, 1, 100) != 29 {
		t.Errorf("count err.")
	}

	if movingCount(10, 1, 10) != 10 {
		t.Errorf("count err.")
	}

	if movingCount(15, 100, 1) != 79 {
		t.Errorf("count err.")
	}

	if movingCount(15, 10, 1) != 10 {
		t.Errorf("count err.")
	}

	if movingCount(15, 1, 1) != 1 {
		t.Errorf("count err.")
	}

	if movingCount(0, 1, 1) != 1 {
		t.Errorf("count err.")
	}

	if movingCount(-10, 10, 10) != 0 {
		t.Errorf("count err.")
	}
}

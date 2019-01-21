package day9

import "testing"

func TestMin(t *testing.T) {
	array1 := []int{3, 4, 5, 1, 2}
	if Min(array1, len(array1)) != 1 {
        t.Errorf("Min num error")
	}
	
	array2 := []int{3, 4, 5, 1, 1, 2}
	if Min(array2, len(array2)) != 1 {
        t.Errorf("Min num error")
	}

	array3 := []int{3, 4, 5, 1, 2, 2}
	if Min(array3, len(array3)) != 1 {
        t.Errorf("Min num error")
	}

	array4 := []int{1, 0, 1, 1, 1}
	if Min(array4, len(array4)) != 0 {
        t.Errorf("Min num error")
	}

	array5 := []int{1, 2, 3, 4, 5}
	if Min(array5, len(array5)) != 1 {
        t.Errorf("Min num error")
	}

	array6 := []int{1}
	if Min(array6, len(array6)) != 1 {
        t.Errorf("Min num error")
	}	
}
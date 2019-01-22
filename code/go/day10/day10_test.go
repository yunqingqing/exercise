package day10

import "testing"

func TestHasPath(t *testing.T) {
	if HasPath("ABTGCFCSJDEH", 3, 4, "BFCE") != true {
		t.Errorf("has path.")
	}

	if HasPath("ABCESFCSADEE", 3, 4, "SEE") != true {
		t.Errorf("has path.")
	}

	if HasPath("ABTGCFCSJDEH", 3, 4, "ABFB") != false {
		t.Errorf("not has path.")
	}

	if HasPath("ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS", 5, 8, "SLHECCEIDEJFGGFIE") != true {
		t.Errorf("has path.")
	}

	if HasPath("ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS", 5, 8, "SGGFIECVAASABCEHJIGQEM") != true {
		t.Errorf("has path.")
	}

	if HasPath("ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS", 5, 8, "SGGFIECVAASABCEEJIGOEM") != false {
		t.Errorf("not has path.")
	}

	if HasPath("AAAAAAAAAAAA", 3, 4, "AAAAAAAAAAAAA") != false {
		t.Errorf("not has path.")
	}
}
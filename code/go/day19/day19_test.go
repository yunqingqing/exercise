package day19

import (
	"testing"
)

func TestFunc(t *testing.T) {
	if IsNumeric("100") != true {
		t.Errorf("not correct.")
	}

	if IsNumeric("123.45e+6") != true {
		t.Errorf("not correct.")
	}

	if IsNumeric("+500") != true {
		t.Errorf("not correct.")
	}

	if IsNumeric("5e2") != true {
		t.Errorf("not correct.")
	}

	if IsNumeric("3.1416") != true {
		t.Errorf("not correct.")
	}

	if IsNumeric("600.") != true {
		t.Errorf("not correct.")
	}

	if IsNumeric("-.123") != true {
		t.Errorf("not correct.")
	}

	if IsNumeric("-1E-16") != true {
		t.Errorf("not correct.")
	}

	if IsNumeric("1.79769313486232E+308") != true {
		t.Errorf("not correct.")
	}

	if IsNumeric("12e") != false {
		t.Errorf("not correct.")
	}

	if IsNumeric("1a3.13") != false {
		t.Errorf("not correct.")
	}

	if IsNumeric("1+23") != false {
		t.Errorf("not correct.")
	}

	if IsNumeric("1.2.3") != false {
		t.Errorf("not correct.")
	}

	if IsNumeric("+-5") != false {
		t.Errorf("not correct.")
	}

	if IsNumeric("12e+5.4") != false {
		t.Errorf("not correct.")
	}

	if IsNumeric(".") != false {
		t.Errorf("not correct.")
	}

	if IsNumeric(".e1") != false {
		t.Errorf("not correct.")
	}

	if IsNumeric("e1") != false {
		t.Errorf("not correct.")
	}

	if IsNumeric("+.") != false {
		t.Errorf("not correct.")
	}

	if IsNumeric("") != false {
		t.Errorf("not correct.")
	}

	// Test("Test21", nullptr, false);
	if IsNumeric("") != false {
		t.Errorf("not correct.")
	}
}

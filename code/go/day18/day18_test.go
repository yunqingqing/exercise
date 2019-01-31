package day18

import "testing"

func TestFunc(t *testing.T) {
	if Math("", "") != true {
		t.Errorf("error.")
	}

	if Math("", ".*") != true {
		t.Errorf("error.")
	}

	if Math("", ".") != false {
		t.Errorf("error.")
	}

	if Math("", "c*") != true {
		t.Errorf("error.")
	}

	if Math("a", ".*") != true {
		t.Errorf("error.")
	}

	if Math("a", "a.") != false {
		t.Errorf("error.")
	}

	if Math("a", "") != false {
		t.Errorf("error.")
	}

	if Math("a", ".") != true {
		t.Errorf("error.")
	}

	if Math("a", "ab*") != true {
		t.Errorf("error.")
	}

	if Math("a", "ab*a") != false {
		t.Errorf("error.")
	}

	if Math("aa", "aa") != true {
		t.Errorf("error.")
	}

	if Math("aa", "a*") != true {
		t.Errorf("error.")
	}

	if Math("aa", ".*") != true {
		t.Errorf("error.")
	}

	if Math("aa", ".") != false {
		t.Errorf("error.")
	}

	if Math("ab", ".*") != true {
		t.Errorf("error.")
	}

	if Math("ab", ".*") != true {
		t.Errorf("error.")
	}

	if Math("aaa", "aa*") != true {
		t.Errorf("error.")
	}

	if Math("aaa", "aa.a") != false {
		t.Errorf("error.")
	}

	if Math("aaa", "a.a") != true {
		t.Errorf("error.")
	}

	if Math("aaa", ".a") != false {
		t.Errorf("error.")
	}

	if Math("aaa", "a*a") != true {
		t.Errorf("error.")
	}

	if Math("aaa", "ab*a") != false {
		t.Errorf("error.")
	}

	if Math("aaa", "ab*ac*a") != true {
		t.Errorf("error.")
	}

	if Math("aaa", "ab*a*c*a") != true {
		t.Errorf("error.")
	}

	if Math("aaa", ".*") != true {
		t.Errorf("error.")
	}

	if Math("aab", "c*a*b") != true {
		t.Errorf("error.")
	}

	if Math("aaca", "ab*a*c*a") != true {
		t.Errorf("error.")
	}

	if Math("aaba", "ab*a*c*a") != false {
		t.Errorf("error.")
	}

	if Math("bbbba", ".*a*a") != true {
		t.Errorf("error.")
	}

	if Math("bcbbabab", ".*a*a") != false {
		t.Errorf("error.")
	}
}

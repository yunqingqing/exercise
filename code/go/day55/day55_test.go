package day55

import "testing"

func TestSearch(t *testing.T) {
	tr := NewTrie()
	tr.Insert("abc")
	tr.Insert("cool")
	tr.Insert("ab")
	tr.Insert("zxb")

	if (tr.Search("abc") != true) {
		t.Errorf("not correct.")
	}
	if (tr.Search("zxb") != true) {
		t.Errorf("not correct.")
	}
	if (tr.Search("z") != false) {
		t.Errorf("not correct.")
	}
	if (tr.Search("a") != false) {
		t.Errorf("not correct.")
	}
}

func TestLookup(t *testing.T) {
	tr := NewTrie()
	tr.Insert("abc")
	tr.Insert("abc")
	tr.Insert("cool")
	tr.Insert("ab")
	tr.Insert("zxb")
	tr.Insert("abc")

	if (tr.Lookup("abc") != 3) {
		t.Errorf("not correct.")
	}
	if (tr.Lookup("zxb") != 1) {
		t.Errorf("not correct.")
	}
	if (tr.Lookup("z") != 0) {
		t.Errorf("not correct.")
	}
}

func TestDelete(t *testing.T) {
	tr := NewTrie()
	tr.Insert("abc")
	tr.Insert("abc")
	tr.Insert("cool")
	tr.Insert("ab")
	tr.Insert("zxb")
	tr.Insert("abc")
	tr.Delete("abc")
	tr.Delete("abc")
	tr.Delete("abc")

	if (tr.Lookup("abc") != 0) {
		t.Errorf("not correct.")
	}
	if (tr.Lookup("zxb") != 1) {
		t.Errorf("not correct.")
	}
	if (tr.Lookup("z") != 0) {
		t.Errorf("not correct.")
	}
}
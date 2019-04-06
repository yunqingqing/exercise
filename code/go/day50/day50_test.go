package day50

import "testing"

func TestFunc(t *testing.T) {
	cache := NewLRUCache(5)
	cache.Put("a", "a")
	if cache.Get("a") != "a" {
		t.Errorf("not corrcet.")
	}

	cache.Put("b", "b")
	if cache.Get("b") != "b" {
		t.Errorf("not corrcet.")
	}

	cache.Put("c", "c")
	if cache.Get("c") != "c" {
		t.Errorf("not corrcet.")
	}

	cache.Put("d", "d")
	if cache.Get("d") != "d" {
		t.Errorf("not corrcet.")
	}

	cache.Put("e", "e")
	if cache.Get("e") != "e" {
		t.Errorf("not corrcet.")
	}

	cache.Put("f", "f")
	if cache.Get("a") != "" {
		t.Errorf("not corrcet. a had been removed.")
	}
	if cache.Get("b") != "b" {
		t.Errorf("not corrcet.")
	}
	if cache.Get("c") != "c" {
		t.Errorf("not corrcet.")
	}
	if cache.Get("d") != "d" {
		t.Errorf("not corrcet.")
	}
	if cache.Get("e") != "e" {
		t.Errorf("not corrcet.")
	}
	if cache.Get("f") != "f" {
		t.Errorf("not corrcet.")
	}

	cache.Put("b", "B")
	if cache.Get("b") != "B" {
		t.Errorf("not corrcet.")
	}

	cache.Put("g", "g")
	if cache.Get("b") != "B" {
		t.Errorf("not corrcet.")
	}
	if cache.Get("c") != "" {
		t.Errorf("not corrcet.")
	}
}

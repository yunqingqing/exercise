package linkedlist

import "testing"

var ll List

func TestAppend(t *testing.T) {
	if !ll.IsEmpty() {
        t.Errorf("list should be empty")
    }
	ll.Append(1)
    if ll.IsEmpty() {
        t.Errorf("list should not be empty")
	}
    if len := ll.Len(); len != 1 {
        t.Errorf("wrong count, expected 1 and got %d", len)
    }
    ll.Append(2)
	ll.Append(4)
    if len := ll.Len(); len != 3 {
        t.Errorf("wrong count, expected 3 and got %d", len)
    }
}
package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(6, 9)
	if result != 15 {
		t.Errorf("Sum was incorrect, want: %d got %d", 15, result)
	}
}

package main

import (
	"log"
	"testing"
)

func TestSubtract(t *testing.T) {
	got := subtract(1, 1)
	want := 0
	if got != want {
		log.Fatalf("we want %v, but we got %v", want, got)
	}
}

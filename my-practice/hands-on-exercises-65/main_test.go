package main

import (
	"log"
	"testing"
)

func TestCar(t *testing.T) {
	got := car("G-Class")
	want := "this is my favorite car G-Class"
	if got != want {
		log.Fatalf("error - want %s, but got %s", want, got)
	}
}

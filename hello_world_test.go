package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Peter")
	want := "Hi,Peter"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
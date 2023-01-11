package main

import "testing"

func TestPrint(t *testing.T) {
	
	testCards := deck{"card 1", "card 2", "card 3", "card 4", "card 5"}

	got := testCards.printCards()
	want :=  " card 1 card 2 card 3 card 4 card 5";

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
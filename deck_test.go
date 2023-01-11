package main

import (
	"reflect"
	"testing"
)

func TestPrint(t *testing.T) {
	
	testCards := deck{"card 1", "card 2", "card 3", "card 4", "card 5"}

	got := testCards.printCards()
	want :=  " card 1 card 2 card 3 card 4 card 5";

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
func TestDeal(t *testing.T) {
	
	cards := deck{"hi", "How", "are", "you"}

	have,got := cards.Deal(3)
	want,value := cards[:3], cards[3:]

	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %q but wanted %q",have, want)
	}
	if !reflect.DeepEqual(got, value) {
		t.Errorf("got %q but wanted %q",got, value)
	}
}
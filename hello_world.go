package main

import "fmt"



func main() {
	newDeck := NewDeck()
	printedOutDeck := newDeck.printCards()
	fmt.Println(printedOutDeck)
	newDeck.SaveToDrive("decks");
}

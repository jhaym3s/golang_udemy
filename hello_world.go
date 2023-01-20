package main

import "fmt"




func main() {
	newDeck := readFromDrive("decks")
	fmt.Println(newDeck.printCards())

	newPerson := Person{name: "Jhaymes", age: 30}

	fmt.Println(newPerson.Birthday())

	fmt.Println(newPerson.pointerAge())
}

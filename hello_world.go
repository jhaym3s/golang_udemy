package main

import "fmt"




func main() {
	newDeck := readFromDrive("deck")
	fmt.Println(newDeck.printCards())
}

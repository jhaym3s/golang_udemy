package main

import "fmt"




func main() {
	newDeck := readFromDrive("decks")
	fmt.Println(newDeck.printCards())

	wallet := Wallet{currency: "Bitcoin",balance: 30}
	wallet.Deposit(50)
	fmt.Printf("wallet sum %d", wallet.sum())
	jhaymes := Person{name: "Paulinus"}
	jhaymes.updateName("James")
	fmt.Println(jhaymes.printName())
}

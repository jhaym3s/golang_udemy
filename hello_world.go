package main

import "fmt"




func main() {
	newDeck := readFromDrive("decks")
	fmt.Println(newDeck.printCards())

	wallet := Wallet{currency: "Bitcoin",balance: 30}
	wallet.Deposit(50)
	fmt.Println(wallet.sum())
}

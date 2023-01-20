package main

import "fmt"
type Wallet struct{
	balance int
	currency string
}

 func (w *Wallet)Deposit(amount int)  {
	w.balance = w.balance+amount
	fmt.Printf("%d %s was deposited ",w.balance, w.currency)
 }

 func (w *Wallet) sum() int {
	return w.balance
 }


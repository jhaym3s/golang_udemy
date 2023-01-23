package main

import "fmt"
type Wallet struct{
	balance int
	currency string
}
type Person struct{
	name string 
}

 func (w *Wallet)Deposit(amount int)  {
	w.balance = w.balance+amount
	fmt.Printf("total balance now is %d",w.balance)
 }

 func (w *Wallet) sum() int {
	return w.balance
 }

 func (p *Person) updateName(newName string)  {
	p.name = newName
 }

 func (p *Person) printName() string{
	return p.name
 }


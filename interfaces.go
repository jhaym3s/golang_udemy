package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bots interface{
	getGreeting() string
}

func (englishBot) getGreeting() string {
	return "hi"
}

func (spanishBot) getGreeting() string {
	return "hi"
}

func printGreeting(b bots)  {
	fmt.Println(b.getGreeting())
}
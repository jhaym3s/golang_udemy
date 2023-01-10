package main

import "fmt"

func main()  {
	fmt.Println(Greet("Jhaymes"))
}

func Greet(name string) string{
	return "Hi,"+name
}
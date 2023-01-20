package main

import "fmt"

type Person struct{
	name string;
	age int;
}

func (p Person)PrintName()  {
fmt.Printf(p.name)
}

func (p *Person) Birthday() int {
	newAge:= p.age+1
	return newAge
}

func (p *Person) pointerAge() int {
	
	return p.age
}
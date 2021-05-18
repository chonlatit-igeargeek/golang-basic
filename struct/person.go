package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

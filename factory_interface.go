package main

import "fmt"

type person struct {
	name     string
	age      int
	position string
}

type Person interface {
	PrintUserData()
}

func (p *person) PrintUserData() {
	fmt.Println("The person is", p)
}

func NewPerson(name string, age int) func(position string) Person {
	return func(p string) Person {
		return &person{name: name, age: age, position: p}
	}
}

func factoryInterface() {
	fmt.Println("Factory with Interface")
	personFactory := NewPerson("NewPerson1", 26)
	gabriel := personFactory("developer")
	gabriel.PrintUserData()
}

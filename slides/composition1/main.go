package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name     string
	Birthday time.Time
}

func (p Person) Say(something string) {
	fmt.Printf("%s: %s\n", p.Name, something)
}

func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

// START OMIT

type Programmer struct {
	Person
}

func (p Programmer) DoProgramming() {
	fmt.Printf("%s is programming now\n", p.Name)
}

func main() {
	ac := Person{
		Name: "AC",
	}
	programmer := Programmer{ac}
	programmer.Say("Hello")
	programmer.DoProgramming()
}

// END OMIT

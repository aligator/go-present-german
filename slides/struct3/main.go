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

// START OMIT

// (p *Person)
func (p Person) ChangeName(newName string) {
	p.Name = newName
}

func main() {
	ac := Person{
		Name: "AC",
	}
	ac.Say("Hello")
	ac.ChangeName("Johannes")
	ac.Say("Hello again")
}

// END OMIT

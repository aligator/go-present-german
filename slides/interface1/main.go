package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Car struct {
	Model string
}

// START OMIT

type Mover interface {
	Move()
}

func (p Person) Move() {
	fmt.Printf("%s walks.\n", p.Name)
}

func (c Car) Move() {
	fmt.Printf("The %s drives.\n", c.Model)
}

func main() {
	var things []Mover = []Mover{
		Person{Name: "AC"},
		Car{Model: "BMW"},
	}

	for _, thing := range things {
		thing.Move()
	}
}

// END OMIT

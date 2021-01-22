package main

import (
	"fmt"
	"time"
)

// START OMIT

type Person struct {
	Name     string
	Birthday time.Time
}

func (p Person) Say(something string) {
	fmt.Printf("%s: %s\n", p.Name, something)
}

func main() {
	ac := Person{
		Name: "AC",
	}

	ac.Say("Hello")
}

// END OMIT

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

func main() {
	ac := Person{
		Name: "AC",
		//Birthday: time.Date(1995, 3, 9, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(ac)
}

// END OMIT

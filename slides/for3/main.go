package main

import "fmt"

// START OMIT

func main() {
	data := []string{"H", "e", "l", "l", "o"}

	var result string

	for i, value := range data {
		fmt.Printf("%v: %v\n", i, value)
		result += value
	}
	fmt.Println(result)
}

// END OMIT

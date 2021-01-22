package main

import (
	"errors"
	"fmt"
)

// START OMIT

func Div(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by 0 not allowed")
	}
	return x / y, nil
}

func main() {
	result, err := Div(42, 2)
	if err != nil {
		fmt.Println("Error while dividing", err)
		return
	}
	fmt.Println(result)

	result, err = Div(42, 0)
	if err != nil {
		fmt.Println("Error while dividing", err)
		return
	}
	fmt.Println(result)
}

// END OMIT

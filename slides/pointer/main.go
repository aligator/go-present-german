package main

import "fmt"

// START OMIT

func main() {
	var aNilPointer *int
	var aNilPointer2 *int = nil

	aVariable := 5
	var aPointer *int = &aVariable
	aVariable++

	fmt.Println(aNilPointer, aNilPointer2, aVariable, aPointer, *aPointer)
}

// END OMIT

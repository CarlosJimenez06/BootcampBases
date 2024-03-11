package main

import "fmt"

func main() {
	// Setting the pointer
	var pointer *int
	// Setting the value
	value := 10
	// Assigning the value to the pointer
	pointer = &value
	// Printing the pointer
	fmt.Println("Original Direction Pointer: ", pointer)
	// Printing the value
	fmt.Println("Value Original", *pointer)
	// Changing the value at original direction
	*pointer = 90
	// Printing the new value
	fmt.Println("Value Modified", *pointer)

	fmt.Println("Value of direction pointer modified: ", pointer)
}

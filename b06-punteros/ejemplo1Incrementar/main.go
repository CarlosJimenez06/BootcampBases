package main

import "fmt"

func Increase(v *int) {
	*v++
}

func main() {

	var v int = 19

	Increase(&v)
	fmt.Println("The value of v now reads", v)
}

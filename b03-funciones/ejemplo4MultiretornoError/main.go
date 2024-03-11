package main

import (
	"errors"
	"fmt"
)

func division(numerator, denominator float64) (result float64, err error) {
	if denominator == 0 {
		return 0, errors.New("Can not divide by zero")
	}
	result = numerator / denominator
	return
}

func main() {
	r, err2 := division(6, 0)
	fmt.Println("Result: ", r)
	fmt.Println(err2)
}

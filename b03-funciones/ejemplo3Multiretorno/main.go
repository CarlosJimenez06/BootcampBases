package main

import "fmt"

func operations(value1, value2 float64) (float64, float64, float64, float64) {

	summatory := value1 + value2
	substraction := value1 - value2
	multiplication := value1 * value2

	var division float64
	if value2 != 0 {
		division = value1 / value2
	}

	return summatory, substraction, multiplication, division
}

func main() {
	sum, sub, mult, div := operations(6, 2)

	fmt.Println("Summatory: \t\t", sum)
	fmt.Println("Substraction: \t\t", sub)
	fmt.Println("Multiplication: \t\t", mult)
	fmt.Println("Division: \t\t", div)

}

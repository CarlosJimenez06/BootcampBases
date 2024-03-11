package main

import "fmt"

func calculateSalary(value float64) (result float64) {

	if value > 50000 {
		result = value - value*0.17
		return
	} else if value > 150000 {
		result = value - value*0.27
		return
	}
	result = value
	return
}

func main() {
	// Calculating salary
	fmt.Println("Your salary is: ", calculateSalary(60000))
}

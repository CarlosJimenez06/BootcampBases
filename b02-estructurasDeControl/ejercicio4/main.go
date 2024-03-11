package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	// Age o Benjamin
	fmt.Println("Age of Benjamin: ", employees["Benjamin"])

	// How much > 21 years old
	count := 0
	for _, v := range employees {
		if v > 21 {
			count++
		}
		//fmt.Println(v)
	}
	fmt.Println("Quantity > 21 years old: ", count)

	// Adding a new Employee

	employees["Federico"] = 25

	// Deleting an Employee

	delete(employees, "Pedro")

	for _, v := range employees {
		fmt.Println(v)
	}

}

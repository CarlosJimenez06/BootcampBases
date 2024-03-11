package main

import "fmt"

func addition(values ...float64) (result float64) {
	//var result float64

	for _, v := range values {
		result += v
		fmt.Println("Result after op: ", result)
	}

	return result
}

func main() {
	fmt.Println("Final value: ", addition(2, 3, 4, 5, 6, 7))

}

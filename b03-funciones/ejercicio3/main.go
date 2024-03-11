package main

import "fmt"

func calculate(quantityMin int, category string) (salary int) {
	switch {
	case category == "A":
		salary = 1000 * quantityMin / 60
	case category == "B":
		salary = 1500*quantityMin/60 + (1000 * 8 * 20 * 0.2)
	case category == "C":
		salary = 3000*quantityMin/60 + (1000 * 8 * 20 * 0.5)
	}
	return
}

func main() {
	fmt.Println("El salario es: ", calculate(60, "A"))
}

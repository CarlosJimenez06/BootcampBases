package main

import "fmt"

const (
	Addition       = "+"
	Substraction   = "-"
	Multiplication = "*"
	Division       = "/"
)

func calculate(value1, value2 float64, operation string) (result float64) {
	switch operation {
	case Addition:
		return value1 + value2
	case Substraction:
		return value1 - value2
	case Multiplication:
		return value1 * value2
	case Division:
		if value2 != 0 {
			return value1 / value2
		}
	}
	return
}

func main() {
	fmt.Println(calculate(6, 2, Addition))
	fmt.Println(calculate(6, 2, Substraction))
	fmt.Println(calculate(6, 2, Multiplication))
	fmt.Println(calculate(6, 2, Division))
}

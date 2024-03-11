package main

import "fmt"

func opSummatory(value1, value2 float64) float64 {
	return value1 + value2
}
func opSubstraction(value1, value2 float64) float64 {
	return value1 - value2
}
func opMultiplication(value1, value2 float64) float64 {
	return value1 * value2
}
func opDivision(value1, value2 float64) float64 {
	if value2 == 0 {
		return 0
	}
	return value1 / value2
}

func calculate(operator string) func(value1, value2 float64) float64 { // Aquí el retorno viene siendo otra función
	switch operator {
	case "Sum":
		return opSummatory
	case "Sub":
		return opSubstraction
	case "Mult":
		return opMultiplication
	case "Div":
		return opDivision
	}
	return nil
}

func main() {

	oper := calculate("Sum")
	r := oper(2, 5)
	fmt.Println(r)
}

package main

import "fmt"

const (
	Addition       = "+"
	Substraction   = "-"
	Multiplication = "*"
	Division       = "/"
)

func opSumatory(value1, value2 float64) float64 {
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

func operationOrchestrator(values []float64,
	operation func(value1, value2 float64) float64) (result float64) { // Aquí el parámetro de entrada es otra función
	for _, v := range values {
		result = operation(result, v)
	}

	return result
}

func calculate(operator string, values ...float64) (result float64) {
	switch operator {
	case Addition:
		return operationOrchestrator(values, opSumatory)
	case Substraction:
		return operationOrchestrator(values, opSubstraction)
	case Multiplication:
		return operationOrchestrator(values, opMultiplication)
	case Division:
		return operationOrchestrator(values, opDivision)
	}
	return
}

func main() {
	fmt.Println("Addition: ", calculate(Addition, 4, 5, 6))
	fmt.Println("Substration: ", calculate(Substraction, 4, 5, 6))
	fmt.Println("Multiplication: ", calculate(Multiplication, 4, 5, 6))
	fmt.Println("Division: ", calculate(Division, 4, 5, 6))
}

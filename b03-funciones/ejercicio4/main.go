package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func opMinimun(grades ...float64) (min float64, err error) {
	//var min float64
	if err != nil {
		return 0, err
	}
	for i, grade := range grades {
		if i == 0 || grade < min {
			min = grade
		}
	}
	return
}
func opAverage(grades ...float64) (avg float64, err error) {
	var sum float64
	if err != nil {
		return 0, err
	}
	for _, grade := range grades {
		sum += grade
	}
	avg = sum / float64(len(grades))
	return
}
func opMaximum(grades ...float64) (max float64, err error) {
	if err != nil {
		return 0, err
	}
	for i, grade := range grades {
		if i == 0 || grade > max {
			max = grade
		}
	}
	return
}

func orchestrator(op string) func(grades ...float64) (float64, error) {

	switch op {
	case minimum:
		return opMinimun
	case average:
		return opAverage
	case maximum:
		return opMaximum
	}
	return nil
}

func main() {
	op := orchestrator(minimum)
	fmt.Println(op(7.8, 9.2, 5.3, 8.1, 7.7))
	op = orchestrator(average)
	fmt.Println(op(7.8, 9.2, 5.3, 8.1, 7.7))
	op = orchestrator(maximum)
	fmt.Println(op(7.8, 9.2, 5.3, 8.1, 7.7))
}

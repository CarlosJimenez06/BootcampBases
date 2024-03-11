package main

import (
	"bases/00-intro/app/lib/calculator"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Red("Hello, world!")
	fmt.Println(calculator.Value)
	result := calculator.Sum(1, 2)
	fmt.Println(result)
}

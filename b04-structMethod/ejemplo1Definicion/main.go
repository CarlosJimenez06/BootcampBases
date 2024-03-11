package main

import "fmt"

type Person struct {
	Name       string
	Gender     string
	Age        int
	Profession string
}

func main() {

	p1 := Person{"Celeste", "F", 25, "Developer"}

	p2 := Person{
		Name:       "Cosito",
		Gender:     "M",
		Age:        30,
		Profession: "Developer",
	}

	// Access to the fields of the struct
	fmt.Println(p1.Name)

	// Change the value of a field
	p2.Age = 40
	fmt.Println(p2.Age)

	// Create a struct without initializing it
	var p3 Person
	p3.Name = "Cosiambiro"
	fmt.Println(p3.Name)
}

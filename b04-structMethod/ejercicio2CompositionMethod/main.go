package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}
type Employee struct {
	ID       int
	Position string
	MyPerson Person
}

func (e Employee) PrintEmoloyee() {
	fmt.Printf("The employee %s has the ID %d and the position %s\n", e.MyPerson.Name, e.ID, e.Position)
}

func main() {
	// Setting the employee
	e := Employee{ID: 1, Position: "Developer", MyPerson: Person{ID: 1, Name: "John", DateOfBirth: "01/01/2000"}}
	// Printing the employee
	e.PrintEmoloyee()
}

package main

import "fmt"

// Struct minimal
type Location struct {
	City    string
	Country string
}

// Struct that aggregates the Location struct
type Company struct {
	Name string
	Location
}

// Method by pointer to migrate the location of the company
func (c *Company) MigrateLocation(newLocation Location) {
	(*c).Location = newLocation
}

// Struct that aggregates the Company struct
type Employee struct {
	Name     string
	Position string
	Company
}

// Method that prints the information of the employee recopilation all information
func (e Employee) Information() {
	fmt.Printf("The employee %s works as %s in %s, %s. Company Name: %s. Company Location: %s",
		e.Name, e.Position, e.City, e.Country, e.Company.Name, e.Company.Location.City)
}

func main() {

	// Setting the employee
	employee := Employee{
		Name:     "John",
		Position: "Developer",
		Company: Company{
			Name: "Google",
			Location: Location{
				City:    "Mountain View",
				Country: "USA",
			},
		},
	}
	// Printing the information of the employee
	fmt.Println("Employee Initial Information: ")
	employee.Information()

	// Migrating the location of the company
	newLocation := Location{
		City:    "San Francisco",
		Country: "USA",
	}
	employee.Company.MigrateLocation(newLocation)

	// Printing the information of the employee after the migration
	fmt.Println("Employee Information After Migration: ")
	employee.Information()

}

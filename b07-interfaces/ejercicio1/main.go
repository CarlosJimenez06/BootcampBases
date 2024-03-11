package main

import "fmt"

type Student struct {
	Name     string
	LastName string
	DNI      int
	Date     int
}

func (s *Student) Info() string {
	return fmt.Sprintf("Name: %s, Last Name: %s, DNI: %d, Date: %d", s.Name, s.LastName, s.DNI, s.Date)
}

func main() {
	s := Student{
		Name:     "Juan",
		LastName: "Perez",
		DNI:      12345678,
		Date:     2000,
	}

	fmt.Println(s.Info())
}

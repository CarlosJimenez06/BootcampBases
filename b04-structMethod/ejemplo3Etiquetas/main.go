package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"-"` // Ignore this field for its sensitive data
	Cellphone string `json:"cellphone"`
	Address   string `json:"address,omitempty"` // If the field is empty, it is not included in the JSON, just print the Key
}

func main() {
	// Initialize a struct
	myPerson := Person{FirstName: "Celeste", LastName: "Gonzalez", Password: "1234", Cellphone: "123456789", Address: ""}

	// Convert the struct to JSON and value the error
	if personaAsJSON, err := json.Marshal(myPerson); err != nil {
		fmt.Println(err)
	} else {
		// Print the JSON
		fmt.Println(string(personaAsJSON))
	}
}

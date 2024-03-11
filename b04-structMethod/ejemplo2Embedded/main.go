package main

import "fmt"

type Preferences struct {
	Foods  string
	Movies string
	Series string
	Animes string
	Sports string
}

type Person struct {
	Name  string
	Age   int
	Likes Preferences
}

func main() {

	// Create a struct initializing, and initializing the embedded struct
	p1 := Person{
		Name: "Celeste",
		Age:  25,
		Likes: Preferences{
			Foods:  "Pizza",
			Movies: "Comedy",
			Series: "Comedy",
			Animes: "Comedy",
			Sports: "Football",
		},
	}

	// Accesing to the fields of the embedded struct
	fmt.Println(p1.Likes.Foods)

	// Create a struct without initializing it
	p3 := Person{}
	p3.Name = "Cosiambiro"
	p3.Age = 30
	// Adding the values of the embedded struct
	p3.Likes = Preferences{Foods: "Pizza", Movies: "Comedy", Series: "Comedy", Animes: "Comedy", Sports: "Football"}
	fmt.Println(p3.Likes)
}

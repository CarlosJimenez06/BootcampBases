package main

import "fmt"

//import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       int
	Description string
	Category    string
}

// Global variable to store the products
var products []Product

func (p Product) Save(prods Product) string {
	products = append(products, prods)
	return "Product saved successfully"
}
func (p Product) GetAll() {
	for _, v := range products {
		fmt.Println(v)
	}
}
func (p Product) GetByID(id int) Product {
	for _, v := range products {
		if v.ID == id {
			fmt.Println(v)
			return v
		}
	}
	return Product{}
}

func main() {
	// Setting products
	p := Product{ID: 1, Name: "Laptop", Price: 1500, Description: "Laptop", Category: "Electronics"}
	p2 := Product{ID: 2, Name: "Mouse", Price: 1500, Description: "Mouse", Category: "Electronics"}

	// Saving products
	fmt.Println(string(p.Save(p)))
	fmt.Println(string(p.Save(p2)))

	// Getting all products
	p.GetAll()

	// Getting a product by ID
	p.GetByID(1)
}

package main

import "fmt"

const (
	SMALL  = "small"
	MEDIUM = "medium"
	LARGE  = "large"
)

type IProduct interface {
	getPrice() float64
}

type Product struct {
	name     string
	price    float64
	tax      float64
	shipping float64
}

func (p *Product) getPrice() string {
	total := p.price + p.price*p.tax
	return fmt.Sprintf("The price of %s is $%.2f", p.name, total)
}

func newSize(size string, values ...interface{}) Product {
	switch size {
	case SMALL:
		return Product{name: values[0].(string), price: values[1].(float64)} //, tax: values[2].(float64)}
	case MEDIUM:
		return Product{name: values[0].(string), price: values[1].(float64), tax: values[2].(float64)}
	case LARGE:
		return Product{name: values[0].(string), price: values[1].(float64), tax: values[2].(float64), shipping: values[3].(float64)}
	}
	return Product{}
}

func main() {
	// Testing SMALL case
	prod := newSize(SMALL, "Laptop", 1000.0)
	fmt.Println(prod.getPrice())

	// Testing MEDIUM case
	prod2 := newSize(MEDIUM, "Laptop", 2000.0, 0.03)
	fmt.Println(prod2.getPrice())

	// Testing LARGE case
	prod3 := newSize(LARGE, "Laptop", 3000.0, 0.06, 2500.0)
	fmt.Println(prod3.getPrice())
}

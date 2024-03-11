package main

import (
	"fmt"
	//"github.com/go-chi/chi/v5"
)

func main() {
	// Declaración básica
	var myInt int
	myInt = 20

	// Declaración múltiple y tipo implícito
	product, price := "MyProduct", 40.6

	// Declaración de una CONSTANTE
	const status = "ok"

	// Casting
	floatNumber := float32(myInt)

	fmt.Println("Hola!! Testeando", product, price)
	fmt.Println("Hola, Testeando", status, floatNumber)
}
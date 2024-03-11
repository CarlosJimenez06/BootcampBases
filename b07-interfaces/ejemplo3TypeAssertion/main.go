package main

import "fmt"

func main() {
	var i interface{} = "hola"

	// Extrae el tipo y compara con el tipo "string", luego asigna a s
	s := i.(string)
	fmt.Println(s)

	// Extrae el tipo y compara con el tipo "string", luego asigna a s, proporcionando un segundo valor booleano
	s, ok := i.(string)
	fmt.Println(s, ok)

	// Extrae el tipo y compara con el tipo "float64", luego asigna a f, EN este caso no es posible, por lo que f tendra el valor 0 y false
	f, ok := i.(float64)
	fmt.Println(f, ok)

}

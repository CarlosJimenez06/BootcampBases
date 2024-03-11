package main

type Person struct {
	FirstName string
	LastName  string
}

// Función que recibe una copia
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Función que recibe un puntero
func (p *Person) ChangeFirstName(newFirstName string) {
	(*p).FirstName = newFirstName
}

// Función que recibe una copia
func (p Person) ChangeLastName(newLastName string) {
	p.LastName = newLastName
}

func main() {

	person := Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	// Changing through Pointer method
	person.ChangeFirstName("Jane")
	person.ChangeLastName("Smith")

	println(person.FullName())

	// ES COMPLETAMENTE NECESARIO PASAR POR PUNTERO SI SE QUIERE MODIFICAR EL VALOR DE UNA ESTRUCTURA
	// AL PASAR POR COPIA SE ESTA CREANDO UNA NUEVA ESTRUCTURA Y NO SE ESTA MODIFICANDO LA ORIGINAL
}

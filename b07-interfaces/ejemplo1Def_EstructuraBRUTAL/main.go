package main

import "fmt"

type Circle struct {
	Radio float64
}

// Method to calculate the area of a circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radio * c.Radio
} // Method to calculate the perimeter of a circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radio
}

/*
// Function to print the details of a circle (PASÁNDOLE EL STRUCT DIRECTAMENTE)

	func details(c Circle) {
		fmt.Println(c)
		fmt.Println("The area of the circle is", c.Area())
		fmt.Println("The perimeter of the circle is", c.Perimeter())
	}
*/
// Function to print the details of a geometry (PASÁNDOLE LA INTERFACE)
func details(g geometry) {
	fmt.Println(g)
	fmt.Println("The area of the geometry is", g.Area())
	fmt.Println("The perimeter of the geometry is", g.Perimeter())
}

/**
* INTERFACE
 */

type geometry interface {
	Area() float64
	Perimeter() float64
}

/**
* Agregando un nuevo tipo de dato
 */
type rect struct {
	width, height float64
}

func (r rect) Area() float64 {
	return r.width * r.height
}
func (r rect) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

/**
* Crea una variable de tipo Circle
 */
/*
func newCircle(value float64) Circle {
	return Circle{
		Radio: value,
	}
}
*/
const (
	rectType   = "RECT"
	circleType = "CIRCLE"
)

// Function to create a new geometry, this returns a geometry interface
func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case rectType:
		return rect{
			width:  values[0],
			height: values[1],
		}
	case circleType:
		return Circle{
			Radio: values[0],
		}
	}
	return nil
}

func main() {
	/* OLD IMPLEMENTATION
	c := Circle{Radio: 10}
	r := rect{width: 10, height: 5}
	c2 := newCircle(4)

	details(c)
	details(r)

	fmt.Println("------------------------------")

	fmt.Println(c2.Area())
	fmt.Println(c2.Perimeter())
	*/

	r := newGeometry(rectType, 10, 5)
	c := newGeometry(circleType, 10)

	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())

	fmt.Println(c.Area())
	fmt.Println(c.Perimeter())
}

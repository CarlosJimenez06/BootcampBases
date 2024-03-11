package main

//import "fmt"

type Circle struct {
	Radio float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

/*
** Method: I can use the same name of the function, but I need to specify the receiver.
** (That behaviour is similar to overwrite a function in Java)
 */

func (c Circle) Area() float64 {
	return 3.14 * c.Radio * c.Radio
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

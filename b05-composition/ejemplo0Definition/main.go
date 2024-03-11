package main

import "fmt"

/**
* Definition of the structs that will be used in the composition
 */

type Engine struct {
	HorsePower int
	Type       string
}
type Chassis struct {
	Material string
}
type Bodywork struct {
	Color string
}

// Interface tha any type of vehicle can implements
type Vehicle interface {
	Start()
	DisplayInfo()
}

/**
* Structs for Auto and Moto that compose the Engine, Chassis and Bodywork structs
 */
type Car struct {
	Engine
	Chassis
	Bodywork
	NumberDoors int
}
type Motorcycle struct {
	Engine
	Chassis
	Bodywork
}

// Methods for the Car struct
func (c Car) Start() {
	fmt.Println("The car is starting")
}
func (c Car) DisplayInfo() {
	fmt.Printf("The car has %d doors, %d horse power, %s type of engine, %s material of chassis and %s color of bodywork\n",
		c.NumberDoors, c.HorsePower, c.Type, c.Material, c.Color)
}

// Methods for the Motorcycle struct
func (m Motorcycle) Start() {
	fmt.Println("The motorcycle is starting")
}
func (m Motorcycle) DisplayInfo() {
	fmt.Printf("The motorcycle has %d horse power, %s type of engine, %s material of chassis and %s color of bodywork\n",
		m.HorsePower, m.Type, m.Material, m.Color)
}

func main() {

	/**
	* Creating an Instance of a Car
	 */
	c := Car{
		Engine: Engine{
			HorsePower: 200,
			Type:       "Gasoline",
		},
		Chassis: Chassis{
			Material: "Steel",
		},
		Bodywork: Bodywork{
			Color: "Red",
		},
		NumberDoors: 4,
	}

	/**
	* Creating an Instance of a Motorcycle
	 */
	m := Motorcycle{
		Engine: Engine{
			HorsePower: 100,
			Type:       "Gasoline",
		},
		Chassis: Chassis{
			Material: "Steel",
		},
		Bodywork: Bodywork{
			Color: "Blue",
		},
	}

	/**
	* Calling the methods of the Car and Motorcycle structs
	 */
	c.Start()
	c.DisplayInfo()

	m.Start()
	m.DisplayInfo()

	/**
	* Calling the implemented method through an Interface
	 */
	var vehicle Vehicle
	vehicle = c
	vehicle.Start()
	vehicle.DisplayInfo()

	vehicle = m
	vehicle.Start()
	vehicle.DisplayInfo()

}

package main

import "fmt"

const (
	TARANTULA = "tarantula"
	DOG       = "dog"
	CAT       = "cat"
	HAMSTER   = "hamster"
)

func CalcTarantulaFood(quantity int) int {
	return 150 * quantity
}
func CalcDogFood(quantity int) int {
	return 10000 * quantity
}
func CalcCatFood(quantity int) int {
	return 5000 * quantity
}
func CalcHamsterFood(quantity int) int {
	return 250 * quantity
}

func Animal(typeAnimal string) (func(int) int, error) {
	switch typeAnimal {
	case TARANTULA:
		return CalcTarantulaFood, nil
	case DOG:
		return CalcDogFood, nil
	case CAT:
		return CalcCatFood, nil
	case HAMSTER:
		return CalcHamsterFood, nil
	}
	return nil, fmt.Errorf("Animal %s not found", typeAnimal)
}

func main() {

	animal, err := Animal(TARANTULA)
	food := animal(2)
	if err != nil {
		return
	}
	fmt.Println("The food for 2 tarantulas is", food, "grams")

	animal, err = Animal(DOG)
	if err != nil {
		return
	}
	food = animal(5)
	fmt.Println("The food for 2 dogs is", animal(2), "grams")
}

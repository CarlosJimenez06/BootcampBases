package main

import "fmt"

func is22major(age int) bool {
	if age > 22 {
		return true
	}
	return false
}
func isEmployee(empl string) bool {
	if empl == "yes" {
		return true
	}
	return false
}
func isAntique(age int) bool {
	if age > 12 {
		return true
	}
	return false
}
func isValue(salary int) bool {
	if salary > 100000 {
		return true
	}
	return false
}

func main() {

	myAge := 45
	amEmployee := "yes"
	myAntique := 14
	myValue := 160000

	switch {
	case is22major(myAge) && isEmployee(amEmployee) && isAntique(myAntique) && isValue(myValue):
		fmt.Println("Your loan havent tax, congrats!!")
	case is22major(myAge) && isEmployee(amEmployee) && isAntique(myAntique):
		fmt.Println("You must pay my taxes, srry!!")
	default:
		fmt.Println("Get out of here!!")
	}

}

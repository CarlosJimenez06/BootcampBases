package main

import "fmt"

func convertionMonth(month int) {
	months := []string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}
	if month > 12 {
		fmt.Println("Non exist that month")
		return
	}
	fmt.Println("Your month is: ", months[month-1])

}

func main() {
	month := 14

	convertionMonth(month)
}

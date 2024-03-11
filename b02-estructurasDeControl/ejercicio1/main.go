package main

import "fmt"

/*
func QLetters(myStr string) int {
	var myInt = myStr.len()
	return
}
*/

func main() {
	var myStr = "abcde"

	fmt.Println("La longitud es: ", len(myStr))

	for i, arr := range myStr {
		fmt.Printf("Letra # %d: %c\n", i, arr)
	}

}

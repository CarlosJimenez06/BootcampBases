package main

import (
	"errors"
	"fmt"
)

func averageGrades(grades ...float64) (average float64, err error) { // Recibir un elipsis para una cantidad INDETERMINADA de valores

	var count float64 = 0
	//var total int = 0;
	for _, v := range grades {
		if v < 0 {
			return 0, errors.New("Negative Grade")
		}
		count = count + v
	}

	average = count / float64(len(grades))
	return
}

func main() {
	avg, err := averageGrades(4, 5, 6) // Necesaria la asignación, pone problema por el múltiple retorno
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Average of Grades: ", avg)

}

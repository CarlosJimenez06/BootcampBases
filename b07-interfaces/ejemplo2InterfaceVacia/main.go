package main

import "fmt"

type SliceAny struct {
	Data []interface{}
}

func main() {
	l := SliceAny{}
	l.Data = append(l.Data, 1, 2)
	l.Data = append(l.Data, "hola")
	l.Data = append(l.Data, true)

	fmt.Printf("%v\n", l.Data)
}

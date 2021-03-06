package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Empty Interface"
	describe(s)
	i := 5
	describe(i)
	strt := struct {
		name string
	}{
		name: "Krish",
	}
	describe(strt)
}

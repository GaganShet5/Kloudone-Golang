package main

import (
	"fmt"
)

func print_hello() {
	fmt.Println("hai")
	print_hello()
}
func main() {
	print_hello()
}

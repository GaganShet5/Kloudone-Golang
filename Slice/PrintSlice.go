package main

import "fmt"

func main() {

	arr := [9]int16{2, 3, 4, 2, 9, 7, 8,,10,4}

	fmt.Println("Array:", arr)

	myslice := arr[7:]
	fmt.Println(myslice)
}

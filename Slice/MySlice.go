package main

import "fmt"

func main() {

	arr := [5]int{2, 3, 7, 6, 9}

	fmt.Println("Array:", arr)

	myslice := arr[0:3]

	fmt.Println("Slice:", myslice)
	fmt.Printf("Length of the slice: %d", len(myslice))
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
}

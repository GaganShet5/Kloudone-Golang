package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("hem")
	fmt.Println(re.FindStringIndex("hellomotto"))
	fmt.Println(re.FindStringIndex("heman"))
	fmt.Println(re.FindStringIndex("washe"))
}

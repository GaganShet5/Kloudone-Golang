package main

import (
	"fmt"
	"strings"
)

func main() {

	s := []string{"gagan", "sullia"}

	for x := 0; x < len(s); x++ {

		res := strings.ToUpper(s[x])

		fmt.Println(res)
	}
}

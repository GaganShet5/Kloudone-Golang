package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading errors", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

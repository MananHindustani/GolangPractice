package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	/*	data, err := ioutil.ReadFile("test.txt")
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		fmt.Println("Contents of file:\n", string(data)) */
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:\n", string(data))
}

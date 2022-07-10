package main

import (
	"fmt"
	"math"
)

const (
	pi = 3.14
)

func main() {
	const fi = 3.15
	a, b := 140.1, 140.01
	c := math.Min(a, b)
	fmt.Println("Minimum value is ", c)
	c = 2 * pi * a

	fmt.Println("CIRCUMFRENCE value is ", c)
	c = 2 * fi * a
	fmt.Println("Ciecumfrence value is ", c)
}

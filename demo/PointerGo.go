package main

import "fmt"

func main() {

	a := 100
	var b *int = &a
	fmt.Printf("Type of a is %T ", b)
	fmt.Printf("\naddress of b is %v", b)
	fmt.Printf("\nType of a is %d ", *b)
	*b++
	b++
	fmt.Printf("\nType of a is %d ", *b)

	fmt.Printf("\nType of a is %d ", a)

	size := new(int)
	fmt.Printf("\nSize value is %d, type is %T, address is %v\n", *size, size, size)
}

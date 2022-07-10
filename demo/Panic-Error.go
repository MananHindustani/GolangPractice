package main

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	firstName := "Manan "
	//	lastName := "Sharma"
	fullName(&firstName, nil)

	//fullName(&firstName, &lastName)
	fmt.Println("returned normally from main")
}

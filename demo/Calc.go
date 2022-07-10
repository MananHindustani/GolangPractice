package main

import "fmt"

func main() {
	var n1 int
	fmt.Println("Enter Number 1")
	fmt.Scanf("%d", &n1)
	var n2 int
	fmt.Println("Enter Number 2")
	fmt.Scanf("%d", &n2)
	var n3 int
	var choice int
	fmt.Println("Enter Choice")
	fmt.Println("1:Add")
	fmt.Println("2:Subtract")
	fmt.Scanf("%d", &choice)
	if choice == 1 {
		n3 = n1 + n2
	} else if n1 > n2 {
		n3 = n1 - n2
	} else {
		n3 = n2 - n1
	}
	fmt.Printf("Ans %d\n", n3)

}

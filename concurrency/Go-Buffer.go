package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	//DEADLOCK	ch <- "steve"
	fmt.Println(<-ch)
	ch <- "steve"

	fmt.Println(<-ch)

}

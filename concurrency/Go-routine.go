package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}
func main() {
	done := make(chan bool)
	go hello(done)
	//	time.Sleep(1 * time.Second)
	<-done
	fmt.Println("main function")
}

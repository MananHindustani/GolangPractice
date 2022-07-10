package main

import "fmt"

func add(x int ,y int )int {
	return  x+y;
}

func main(){
	var z=add(4, 1);
	fmt.Println("Add ",z)

	fmt.Println( add(1,5));

}
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	//	fmt.Println(array);
	//	fmt.Println(array[1])
	//	fmt.Printf("%d\n",array)
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a)
	fmt.Printf("%s\n", a)
	fmt.Printf("%q\n", a)
	var array []int
	fmt.Println("Enter Element for Array")
	for i := 0; i < 4; i++ {
		fmt.Scanf("%d", &array[i])
	}
	for v := range array {
		fmt.Printf("%d", v)
	}
	array1 := array

	fmt.Println(array1)
	fmt.Println(array[2:])
	array.append(array, 1)
	fmt.Println(array[2:])
}

	package main

import "fmt"

func main() {
	sum := 0
	aa:
	for i := 1; i <= 10; i++ {
		bb:
		for  j := 1; j <= 10; j++  {
			if (i == 5 && j == 5) {
				continue aa
			} else if i == 10 {
				break bb
			}
		}
		sum += i
		fmt.Println("Hello");

	}
	fmt.Println("sum is", sum);
}
package main

import "fmt"

func addition(a int, b int) int {
	return a + b
}
func substraction(a int, b int) int {
	return a - b
}
func multiply(a int, b int) int {
	return a * b
}
func divide(a float64, b float64) float64 {
	return a / b
}
func swap(a int, b int)  (int,int) {
temp:=0
temp=a
a=b
b=temp
return a, b
}


func main() {
	var n1 int
	fmt.Println("Enter Number 1")
	fmt.Scanf("%d", &n1)
	var n2 int
	fmt.Println("Enter Number 2")
	fmt.Scanf("%d", &n2)
	var choice int
	fmt.Println("Enter choices")
	fmt.Println("1:Add")
	fmt.Println("2:Subtraction")
	fmt.Println("3:Multiply")
	fmt.Println("4:Divide")
	fmt.Println("5:Swap")

	fmt.Scanf("%d", &choice)	case 4:
	fmt.Println("Ans is ", divide(float64(n1), float64(n2)))


	switch choice {
	case 1:
		fmt.Println("Ans is ", addition(n1, n2))
	case 2:
		if temp := 0; n2 > n1 {
			temp = n1
			n1 = n2
			n2 = temp

		}
		fmt.Println("Ans is ", substraction(n1, n2))

	case 3:
		fmt.Println("Ans is", multiply(n1, n2))
	case 4:
		fmt.Println("Ans is ", divide(float64(n1), float64(n2)))
	case 5:
	fmt.Println("Ans is ", divide(float64(n1), float64(n2)))

	default:
		fmt.Println("Wrong Input")
	}

}

package main

import "fmt"

func digits(number int, dgchnl chan int) {

	for number != 0 {
		digit := number % 10
		dgchnl <- digit
		number /= 10
	}
	close(dgchnl)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dgchnl := make(chan int)
	go digits(number, dgchnl)

	for digit := range dgchnl {
		sum = sum + digit*digit
		number /= 10
	}
	/*	for number != 0 {
			digit := number % 10
			sum = sum + digit*digit
			number /= 10
		}
	*/
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dgchnl := make(chan int)
	go digits(number, dgchnl)
	for digit := range dgchnl {
		sum = sum + digit*digit*digit
		number /= 10
	}
	/*	for number != 0 {
			digit := number % 10
			sum = sum + digit*digit*digit
			number /= 10
		}
	*/
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

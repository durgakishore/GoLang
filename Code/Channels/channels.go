package main

import "fmt"

func calcSquares(number int, sqch chan int) {

	sum := 0

	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	sqch <- sum
}
func calcCubes(number int, cubch chan int) {

	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubch <- sum
}
func main() {

	sqch := make(chan int)
	cubch := make(chan int)

	go calcSquares(122, sqch)
	go calcCubes(122, cubch)

	sq, cub := <-sqch, <-cubch

	fmt.Println(sq, cub)
}

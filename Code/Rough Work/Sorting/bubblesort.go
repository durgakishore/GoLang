package main

import "fmt"

func bubbleSort(intergers []int) {

	var isSwapped bool = true
	num := len(intergers)

	for isSwapped {

		isSwapped = false

		for i := 1; i < num; i++ {

			if intergers[i-1] > intergers[i] {
				temp := intergers[i-1]
				intergers[i-1] = intergers[i]
				intergers[i] = temp
				isSwapped = true
			}
		}
	}
	fmt.Println(intergers)
}

func main() {

	numbers := []int{23, 45, 34, 22, 67, 54, 1, 1, 2, 37, 1}
	bubbleSort(numbers)

	str := "kishore"

	for _, s := range str {
		defer fmt.Printf("%c", s)
	}
}

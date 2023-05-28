package main

import "fmt"

func SquareSum(numbers []int) int {
    // your code here
	var square int 
	var sum int
	for i := 0; i < len(numbers); i++ {
		square = numbers[i]*numbers[i]
		sum += square
	}
	return sum
}

func main()  {
	number := []int{2,4}
	fmt.Println(SquareSum(number))
}
package main

import (
	"fmt"
	"strconv"
)


func fizzBuzz(n int) []string {
	var stringOut []string 
	for i := 1; i <= n; i++ {
		answer := ""
		if i%15 == 0 {
			answer = "FizzBuzz"
		} else if i%3 == 0 {
			answer = "Fizz"
		} else if i%5 == 0 {
			answer = "Buzz"
		} else {
			answer = strconv.Itoa(i)
		}
		stringOut = append(stringOut, answer)
	}
	return stringOut
}

func main()  {
	fmt.Println(fizzBuzz(15))
}

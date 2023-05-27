package main

import "fmt"

func Between(a, b int) []int {
	// your code here
	var sum int
	var sumAppend []int
	for i := a; i <= b; i++ {
		sum = i

		sumAppend = append(sumAppend, sum)
	}
	return sumAppend
}

func main()  {
	a := 1
	b := 4
	fmt.Println(Between(a,b))	
}
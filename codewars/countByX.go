package main

import "fmt"

func CountBy(x, n int) []int {
	var row []int
	var sum int
	for i := 0; i < n; i++ {
		sum += x
		row = append(row, sum)
	}
	return row
}

func main()  {
	x := 4
	n := 20
	fmt.Println(CountBy(x,n))
}
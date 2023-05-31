package main

import "fmt"

func Grow(arr []int) int{
	var sum int = 1
	for i := 0; i < len(arr); i++ {
		sum *= arr[i]
	}
	return sum
}

func main()  {
	arr := []int{1,2,3,4}
	fmt.Println(Grow(arr))
}
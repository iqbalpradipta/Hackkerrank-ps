package main

import (
	"fmt"
	"sort"
)

func findMedian(arr []int32) int32 {
    // Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	n := len(arr)
	median := n / 2 
	if n % 2 == 0{
		return arr[median-1] + arr[median] / 2
	}
	return arr[median]
}

func main()  {
	arr := []int32{1,2,6,5,3}
	fmt.Println(findMedian(arr))
}
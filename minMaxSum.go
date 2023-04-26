package main

import (
	"fmt"
	"sort"
	"strconv"
)

func miniMaxSum(arr []int32) {
    // Write your code here
	var max int64
	var min int64
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 0; i < len(arr)-1; i++ {
		min += int64(arr[i])
	} 
	for j := 1; j < len(arr); j++ {
		max += int64(arr[j])
	}
	fmt.Println(strconv.Itoa(int(min)) + " "+ strconv.Itoa(int(max)))
	
}

// func main()  {
// 	arr := []int32{1,2,3,4,5}
// 	miniMaxSum(arr)

// }
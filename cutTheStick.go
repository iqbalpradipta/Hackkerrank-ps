package main

import (
	"fmt"
	"sort"
)

func cutTheSticks(arr []int32) []int32 {
    // Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var cutStick []int32
	for i := 0; i < len(arr); {
		count := 1
		for j := i + 1; j < len(arr) && arr[j] == arr[i]; j++ {
			count++
		}

		cutStick = append(cutStick, int32(len(arr)-i))
		i += count
	}
	return cutStick
}

func main()  {
	arr := []int32{5, 4, 4, 2, 2, 8}
	fmt.Println(cutTheSticks(arr))	
}
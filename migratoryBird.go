package main

import (
	"fmt"
)

func migratoryBirds(arr []int32) int32 {
    // Write your code here
	var sum int32
	for i := 0; i < len(arr); i++ {
		for j := len(arr)-1; j >= 0; j-- {
			if arr[i] == arr[j] {
				sum++
			}
		}
	}
	return sum
}

func main()  {
	arr := []int32{1, 4, 4, 4, 5, 3}

	fmt.Println(migratoryBirds(arr))
}

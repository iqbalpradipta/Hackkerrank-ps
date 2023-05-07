package main

import (
	"fmt"
)

func hurdleRace(k int32, height []int32) int32 {
    // Write your code here
	var max int32 = height[0]
	var result int32

	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}

	if k >= max {
		return 0
	} else {
		result = max - k
		return result
	}
}

func main()  {
	var k int32 = 4
	var height []int32 = []int32{1,6,3,5,2}
	fmt.Println(hurdleRace(k, height))
}

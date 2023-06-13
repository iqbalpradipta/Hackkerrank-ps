package main

import (
	"fmt"
)


func jumpingOnClouds(c []int32) int32 {
    // Write your code here
	var res int32
	for i := 2; i < len(c) - 1; i++ {
		if c[i] == 0 {
			i++
		}
		res++
	}
	return res + 1
}

func main()  {
	c := []int32{0,0,0,0,1,0}
	fmt.Println(jumpingOnClouds(c))
}

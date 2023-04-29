package main

import (
	"fmt"
)

func migratoryBirds(arr []int32) int32 {
    // Write your code here
	bird := map[int32]int32{}
	var count int32
	var res int32
	
	for _, v := range arr {
		bird[v]++ 
		if bird[v] > count {
			res = v
			count = bird[v]
		} else if bird[v] == count && v < res { 
			res = v 	
		}
		
	}
	return res
}

func main()  {
	arr := []int32{1, 4, 4, 4, 5, 3}

	fmt.Println(migratoryBirds(arr))
}

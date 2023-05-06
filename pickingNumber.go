package main

import (
	"fmt"
	"sort"
)

func pickingNumbers(a []int32) int32 {
    // Write your code here
	sort.Slice(a, func(i, j int) bool {
		return a[i] >= a[j]
	})
	var diff int32 = 0
	for i := 0; i < len(a)-1; i++ {
		var count int32 = 0
		for j := i; j < len(a); j++ {
			if a[i] - a[j] <= 1{
				count++
				if count > diff {
					diff =	count
				}
			} else {
				count = 0
			}
		}
	}
	return diff
}

func main()  {
	var a []int32 = []int32{1, 2, 2, 3, 1, 2}
	fmt.Println(pickingNumbers(a))

}
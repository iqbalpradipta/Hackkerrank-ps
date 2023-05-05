package main

import "fmt"

func pickingNumbers(a []int32) int32 {
    // Write your code here

	var diff int32 
	for _, v := range a {
		var count int32
		for _, n := range a{
			if (a[v] - a[n]) <= 1 && a[v] >= a[n] {
				count += 1
				n += 1
			}
		}
		if count > diff {
			diff = count-1
		}
	}
	return diff
}

func main()  {
	var a []int32 = []int32{1, 2, 2, 3, 1, 2}
	fmt.Println(pickingNumbers(a))

}
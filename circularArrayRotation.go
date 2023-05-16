package main

import "fmt"

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
    // Write your code here
	rotasi := make([]int32, len(queries))
	for i := 0; i < int(k); i++ {
		temp := a[len(a)-1]
		for j := len(a)-1; j > 0; j-- {
			a[j] = a[j-1]
		}
		a[0] = temp
	}
	for q, v := range queries {
		rotasi[q] = a[v]
	}
	return rotasi
}

func main()  {

	a := []int32{1,2,3,4,5,6,7,8,9}
	k := 3
	queris := []int32{2}

	fmt.Println(circularArrayRotation(a,int32(k),queris))
}
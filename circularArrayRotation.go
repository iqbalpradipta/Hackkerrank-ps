package main

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
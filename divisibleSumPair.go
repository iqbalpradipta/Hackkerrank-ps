package main

import "fmt"

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    // Write your code here
	var count int32
	for i := 0; int32(i) < n-1; i++ {
		for j := i+1; int32(j) < n; j++ {
			if (ar[i] + ar[j]) % k == 0 {
				count++
			}
		}	
	}
	return count
}

func main()  {
	var n int32 = 6
	var k int32 = 3
	var ar []int32 = []int32{1, 3, 2, 6, 1, 2}
	fmt.Println(divisibleSumPairs(n,k,ar))
}
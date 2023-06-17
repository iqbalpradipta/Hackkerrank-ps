package main

import "fmt"

/*
N = harga coklat
C = dompet
M = gratisan
*/


func chocolateFeast(n int32, c int32, m int32) int32 {
    // Write your code here
	var chocolate int32
	var exchange int32
	chocolate = n/c
	exchange = chocolate
	for exchange >= m {
		exchange = exchange - m + 1
		chocolate++
	}
	return chocolate
}

func main()  {
	n := 6
	c := 2
	m := 2
	fmt.Println(chocolateFeast(int32(n),int32(c),int32(m)))
}
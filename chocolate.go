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
	var result int32

	if chocolate == m {
		result = exchange + 1
	} else if chocolate < m {
		result = exchange
	} else if chocolate > m {
		result = exchange * m - 1
	}
	return result
}

func main()  {
	n := 10
	c := 2
	m := 2
	fmt.Println(chocolateFeast(int32(n),int32(c),int32(m)))
}
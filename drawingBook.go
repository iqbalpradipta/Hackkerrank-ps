package main

import "fmt"

func pageCount(n int32, p int32) int32 {
    // Write your code here
	var front int32 = p/2
	var back int32 = (n-p)/2

	if p % 2 != 0 && n-p == 1 {
		back += 1
	}
	
	if back < front {
		return back
	}
	return front

}

func main()  {
	var n int32 = 5
	var p int32 = 4
	fmt.Println(pageCount(n,p))
}
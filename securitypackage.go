package main

import (
	"fmt"
	"strconv"
)

// Solution is your solution code.
func Solution (x int) []string {
	// Your code starts here.
	var kataOut []string
	for i := 1; i <= x; i++ {
		kata := ""
		if i % 15 == 0 {
			kata = "scan inspect"
		} else if i % 5 == 0 {
			kata = "inspect"
		} else if i % 3 == 0 {
			kata = "scan"
		} else {
			kata = strconv.Itoa(i)
		}
		kataOut = append(kataOut, kata)
	}
	return kataOut
}

func main()  {
	x := 3
	fmt.Println(Solution(x))
}
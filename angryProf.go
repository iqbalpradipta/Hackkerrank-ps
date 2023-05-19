package main

import "fmt"

func angryProfessor(k int32, a []int32) string {
    // Write your code here
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			k--
		}
	}
	
	if k > 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func main()  {
	k := 3
	a := []int32{0 ,-1, 2, 1}
	fmt.Println(angryProfessor(int32(k),a))
}
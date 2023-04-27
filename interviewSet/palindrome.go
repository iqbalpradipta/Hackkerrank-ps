package main

import (
	"fmt"
)

// func isPalindrome(value string) bool {
// 	var sumString string
// 	for i := len(value)-1; i >= 0; i-- {
// 		sumString += string(value[i])
// 	}
// 	fmt.Println(sumString)
// 	return sumString == value
// }

func isPalindromewithoutNewVariable(value string) bool {
	for i := 0; i < len(value); i++ {
		iAwal := i
		iAkhir := len(value)-i-1
		if string(value[iAwal]) != string(value[iAkhir]) {
			return false
		}
	}
	fmt.Println(value)
	return true
}

func main()  {
	fmt.Println(isPalindromewithoutNewVariable("kodok"))
}
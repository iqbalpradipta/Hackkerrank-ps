package main

import "fmt"

func Solution (x string) bool {
	// Your code starts here.
	parenthes := []rune{}

	for _, v := range x {
		if v == '(' || v == '[' || v == '{' {
			parenthes = append(parenthes, v)
		} else if  v == ')' || v == ']' || v == '}' {
			if len(parenthes) == 0 {
				return false
			}
			switch v {
			case ')':
				if parenthes[len(parenthes)-1] != '(' {
					return false
				}
			case ']':
				if parenthes[len(parenthes)-1] != '[' {
					return false
				}
			case '{' :
				if parenthes[len(parenthes)-1] != '{' {
					return false
				}
			}
			parenthes = parenthes[:len(parenthes)-1]
		}
	}
	if len(parenthes) > 0 {
		return false
	}
	return true
}

func main()  {
	x := "[] {} []"
	fmt.Println(Solution(x))
}
package main

import "fmt"

func repeatedString(s string, n int64) int64 {
    // Write your code here
	var count int64
	var res int64

	for i := range s {
		if s[i] != 'a' {
			count++
		}
		if int64(i) < n % int64(len(s)) && s[i] == 'a' {
			res++
		}
	}

	res += n / int64(len(s)) * ((int64(len(s)) - int64(count)))
	return res
}

func main()  {
	s := "aba"
	n := 10
	fmt.Println(repeatedString(s,int64(n)))
}
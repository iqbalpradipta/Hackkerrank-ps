package main

import "fmt"


func saveThePrisoner(n int32, m int32, s int32) int32 {
    // Write your code here
	prisuner := (s-1+m) % n 
	if prisuner == 0 {
		prisuner = n
	}
	return prisuner
}

func main()  {
	var n int32 = 7
	var m int32 = 19
	var s int32 = 2
	fmt.Println(saveThePrisoner(n,m,s))
}
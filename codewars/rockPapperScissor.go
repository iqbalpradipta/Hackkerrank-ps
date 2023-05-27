package main

import "fmt"

func Rps(p1, p2 string) string {
	r := "rock"
	p := "paper"
	s := "scissors"
	var result string


	if p1 == s && p2 == p || p1 == r && p2 == s || p1 == p && p2 == r {
		result = "Player 1 won!"
	} else if p1 == s && p2 == r || p1 == p && p2 == s || p1 == r && p2 == p {
		result = "Player 2 won!"
	} else {
		result = "Draw!"
	}

	return result
}

func main()  {
	p1 := "scissors"
	p2 := "paper"
	fmt.Println(Rps(p1,p2))
}
package main

import (
	// "fmt"
	"time"
)

func timeConversion(s string) string {
    // Write your code here
	t, _ := time.Parse("03:04:05PM", s)
	return t.Format("15:04:05")

}

// func main()  {
// 	s := "07:05:45PM"
// 	fmt.Println(timeConversion(s))
// }
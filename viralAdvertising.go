package main

import (
	"fmt"
	"math"
)


func viralAdvertising(n int32) int32 {
    // Write your code here
    var sum int32
    var berbagi int32 = 5
    for i := 0; int32(i) < n; i++ {
        suka := math.Floor(float64(berbagi)/2)
        berbagi = int32(suka * 3)
        sum += int32(suka) 
        
    }
    return sum
}

func main()  {
    n := 4
    fmt.Println(viralAdvertising(int32(n)))
}
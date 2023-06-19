package main

import (
	"fmt"
	"math"
)

func squares(a int32, b int32) int32 {
    // Write your code here
    var sqrtA float64 = math.Sqrt(float64(a))
    var sqrtB float64 = math.Sqrt(float64(b))
    var count int32
    for i := int32(sqrtA) ; i <= int32(sqrtB) ; i++ {
        if i*i >= a && i*i <= b {
            count ++
        }
    }
    return count
}

func main()  {
	a := 3
	b := 9
	fmt.Println(squares(int32(a), int32(b)))
}
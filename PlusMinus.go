package main

import "fmt"

func plusMinus(arr []int32) {
    // Write your code here
    var resultPositive float64 = 0.0
    var resultNegative float64 = 0.0
    var resultZero float64 = 0.0
    var n float64 = float64(len(arr))
    
    for v := range arr {
        if arr[v] > 0 {
            resultPositive++
        } else if arr[v] < 0 {
            resultNegative++
        } else {
            resultZero++

        }
    }
    fmt.Printf("%.6f\n", resultPositive/n)
    fmt.Printf("%.6f\n", resultNegative/n)
    fmt.Printf("%.6f\n", resultZero/n)
}

func main()  {
	var arr []int32 = []int32{-4, 3, -9, 0, 4, 1}
	// var n int32 = int32(len(arr))

	plusMinus(arr)
}
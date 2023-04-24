package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
    // Write your code here
	var sumLeft int32
	var sumRight int32
	var l int = len(arr)
	for i := 0; i < l; i++ {
		sumLeft += arr[i][i]
		sumRight += arr[i][l-1-i]
		fmt.Println("sumLeft = ", sumLeft)
		fmt.Println("sumRight = ", sumRight)
	}
	result := math.Abs(float64(sumLeft)-float64(sumRight))
	return int32(result)
}

func main (){
	arr := [][]int32{
		{11,2,4},
		{4,5,6},
		{10,8,-12}}
	fmt.Println(diagonalDifference(arr))
}
package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
    // Write your code here
	var sum int32
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			sum += arr[i][j]
			fmt.Println(sum)
		}
	}
	return sum
}

func main (){
	arr := [][]int32{
		{11,2,4},
		{4,5,6},
		{10,8,-12}}
	fmt.Println(diagonalDifference(arr))
}
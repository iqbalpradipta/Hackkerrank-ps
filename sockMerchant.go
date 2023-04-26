package main

// import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
    // Write your code here
	pairs := make(map[int32]int32)
	var res int32
	var finalres int32
	for _, v := range ar {
		pairs[v]++
		if res = pairs[v]; res % 2 == 0 { 
			finalres++
		}
	}
	return finalres
}

// func main()  {
// 	var n int32 = 7
// 	var ar []int32 = []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
// 	fmt.Println(sockMerchant(n, ar))
// }
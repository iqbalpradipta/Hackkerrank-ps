package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32){
    // Write your code here
	var sum int32
	var b2 int32
	var sumsum int32
	for i := 0; i < len(bill); i++ {
		sum += bill[i]
		if int32(i) == k {
			sum -= bill[k]
			b2 = sum % 2
			sumsum = b - b2
		} else if k == b {
			fmt.Println("Bon Appetit")
		}
	} 
	fmt.Println(sumsum)
}

func main()  {
	var bill = []int32{3, 10, 2, 9}
	var k int32 = 1
	var b int32 = 12
	bonAppetit(bill, k, b)
}
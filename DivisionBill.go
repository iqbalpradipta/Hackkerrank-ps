package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32){
    // Write your code here
    var sum int32
    var newBill int32 = k
    bill = append(bill[:newBill], bill[newBill + 1:]...)
    var b2 int32
    var res int32
    var text string = "Bon Appetit"
    for i := 0; i < len(bill); i++ {
        sum += bill[i] 
    }
    b2 = sum / 2
	res = b - b2
	if b2 == b {
		fmt.Println(text)
	} else {
		fmt.Println(res)
	}
}

// func main()  {
// 	var bill = []int32{72, 53, 60, 66, 90, 62, 12, 31, 36, 94}
// 	var k int32 = 6
// 	var b int32 = 288
// 	bonAppetit(bill, k, b)
// }
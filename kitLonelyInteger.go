package main

// import "fmt"



func lonelyinteger(a []int32) int32 {
    // Write your code here
	var sum int32
	for i := 0; i < len(a); i++ {
		sum ^= a[i]
	}
	return sum
}

// func main()  {
// 	a := []int32{0,0,2,4,4,1,1}
// 	fmt.Println(lonelyinteger(a))
// }
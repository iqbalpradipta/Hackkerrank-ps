package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
    /*
     * Write your code here.
     */
    var maxBill int32 = -1

     for i := 0; i < len(keyboards); i++ {
        for j := 0; j < len(drives); j++ {
            total := keyboards[i] + drives[j]
            if total <= b && total >= maxBill{
                maxBill = total
            }
        }
     }
     return maxBill
}

func main()  {
    var keyboards []int32 = []int32{3,1}
    var drives []int32 = []int32{5,2,8} 
    var b int32 = 10
    fmt.Println(getMoneySpent(keyboards, drives, b))
}
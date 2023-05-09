package main

import "fmt"


func beautifulDays(i int32, j int32, k int32) int32 {
    var count int32
    for day := i; day <= j; day++ {
        reversed := reverse(day)
        if abs(day - reversed) % k == 0 {
            count++
        }
    }
    return count
}

func reverse(n int32) int32 {
    var rev int32
    for n > 0 {
        digit := n % 10
        rev = rev*10 + digit
        n /= 10
    }
    return rev
}

func abs(n int32) int32 {
    if n < 0 {
        return -n
    }
    return n
}

func main()  {
    i := int32(20)
    j := int32(23)
    k := int32(6)
    fmt.Println(beautifulDays(i,j,k))
}
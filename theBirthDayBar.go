package main

import "fmt"

/*Rules
1. d panjang coklat yang ingin diberi
2. m itu angka awal yang ingin mendekati d
3. s angka yang ada di coklat
4. m harus di jumlahkan yg sesuai dengan d
*/

func birthday(s []int32, d int32, m int32) int32 {
    // Write your code here 
    var sum int32
    var count int32
    for v := 0; v < len(s); v++ {
        sum += s[v]
        if int32(v) >= m - 1 {
            if sum == d {
                count++
            }
            sum -= s[int32(v)-m+1]
        }
    }
    return count  
}

func main()  {
    var s = []int32{1,2,1,3,2}
    var d int32= 3
    var m int32= 2
    fmt.Println(birthday(s,d,m))
}
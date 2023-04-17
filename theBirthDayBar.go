package main

/*Rules
1. d panjang coklat yang ingin diberi
2. m itu angka awal yang ingin mendekati d
3. s angka yang ada di coklat
4. m harus di jumlahkan yg sesuai dengan d
*/

func birthday(s []int32, d int32, m int32) int32 {
    // Write your code here
    var count int32
	var sum int32
    for i := 0; int32(i) < m; i++ {
        sum += s[i]
    }
    if sum == d {
        count ++
    }
    for i := m; i < int32(len(s)); i++ {
        for j := 0; j < len(s); j++ {
            sum += s[i] + s[j]
            if sum == d {
                count ++
            }
        }
    }
    return count
}
package main

/*Rules
1. d panjang coklat yang ingin diberi
2. m itu angka awal yang ingin mendekati d
3. s angka yang ada di coklat
4. m harus di jumlahkan yg sesuai dengan d
*/

func birthday(s []int32, d int32, m int32) int32 {
    // Write your code here
	for _, v := range s {   
        for _, j := range s {
            if v + j == d {
                return j
            } else if m < v {
                return m
            }
        }
    }
    return 0
}
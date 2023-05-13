package main

import "fmt"

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
    // Write your code here
	tidakTerlambat := 0
	terlambatHari := 15 * (d1 - d2)
	terlambatBulan := 500 * (m1 - m2)
	terlambatTahun := 10000

	var res int32

	if d1 == d2 && m1 == m2 && y1 == y2 {
		return int32(tidakTerlambat)
	} else if d2 < d1 && m1 == m2 && y1 == y2 {
		res = terlambatHari
		return res	
	} else if m2 < m1 && y1 == y2 {
		res = terlambatBulan
		return res
	} else if y2 < y1 {
		res = int32(terlambatTahun)
		return res
	}
	return res
}

func main()  {
	var d1, m1, y1 int32 = 9 , 6, 2015
	var d2, m2, y2 int32 = 6 , 6, 2015
	fmt.Println(libraryFine(d1, m1, y1, d2, m2 ,y2))
}
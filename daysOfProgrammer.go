package main

import (
	"strconv"
	// "fmt"
)

func dayOfProgrammer(year int32) string {
    // Write your code here
	if year > 1918 {
		return kalenderGregorian(year)
	} else if year <= 1917 {
		return KalenderJulian(year)
	} else {
		return "26.09." + strconv.Itoa(int(year))
	}
}

func kalenderGregorian(year int32) string {
    str := strconv.Itoa(int(year))
    if year % 400 == 0 || year % 4 == 0 && year % 100 != 0 {
        return "12.09." + str
    } else {	
        return "13.09." + str
    }
}

func KalenderJulian(year int32) string {

	if year % 4 == 0 {
		return "12.09." + strconv.Itoa(int(year))
	} else {
		return "13.09." + strconv.Itoa(int(year))
	}
}

// func main()  {
// 	fmt.Println(dayOfProgrammer(2019))
// 	fmt.Println(dayOfProgrammer(1918))
// 	fmt.Println(dayOfProgrammer(1700))
// }
package main

import (
	"fmt"
	"math"
)

func catAndMouse(x int32, y int32, z int32) string {
	kucingA := math.Abs(float64(z - x))	
	kucingB := math.Abs(float64(z - y))

	if kucingA < kucingB {
		return "Cat A"
	} else if kucingA > kucingB {
		return "Cat B"
	} else {
		return "Mouse C"
	}
}

func main()  {
	x := 1
	y := 3
	z := 2
	fmt.Println(catAndMouse(int32(x),int32(y),int32(z)))
}
package main

import (
	"fmt"
	"sort"
)

func SmallestIntegerFinder(numbers []int) int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	
	var smallestInt int
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				smallestInt = numbers[i]
				return smallestInt
			} else {
				smallestInt = numbers[j]
				return smallestInt
			}
		}
	}
	return smallestInt
}

func main()  {
	number := []int{34, 15, 88, 2}
	fmt.Println(SmallestIntegerFinder(number))
}
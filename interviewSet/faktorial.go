package main

import "fmt"

func factorial(value int) int {
	if value <= 0 {
		return 1
	}

	result := 1
	for i := value; i >= 1; i-- {
		if i % 2 == 1 {
			result *= i
			fmt.Println("hasil = ", i)
		}
	}
	return result
}

func faktorialRecursive(value int) int {
	if value <= 0 {
		return 1
	} else {
		return value * faktorialRecursive(value - 1)
	}
}

func faktorialTailRecursive(total, value int) int {
	if value <= 0 {
		return total
	} else  {
		return faktorialTailRecursive(total * value, value -1)
	}
}

func main()  {
	fmt.Println(factorial(5))	
	fmt.Println(faktorialRecursive(5))
	fmt.Println(faktorialTailRecursive(1,5))	
}
package main

import (
	"fmt"
	"math"
)

func ListSquared(m, n int) [][]int {
    // your code
	result := [][]int{}

	for number := m; number <= n; number++ {
		sumOfSquare := []int{}

		for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
			if number % i == 0 {
				sumOfSquare = append(sumOfSquare, i)
				if i != number/i {
					sumOfSquare = append(sumOfSquare, number / i)
				}
			} 
		}
		sumNumber := 0 
		for _, v := range sumOfSquare {
			sumNumber += v * v
		}

		if math.Sqrt(float64(sumNumber)) == float64(int(math.Sqrt(float64(sumNumber)))) {
			result = append(result, []int{number, sumNumber})
		}
	}
	return result

}

func main()  {
	m := 1
	n := 250
	fmt.Println(ListSquared(m,n))
}
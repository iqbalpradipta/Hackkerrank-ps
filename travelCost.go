package main

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

func SolutionTravel (input [][]int) int {
	// Your code starts here.
	m := len(input)
	n := len(input[0])

	for i := 1; i < m; i++ {
		input[i][0] +=input[i-1][0]
	}

	for j := 1; j < n; j++ {
		input[0][j] += input[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			input[i][j] += min(input[i-1][j], input[i][j-1])
		}
	}

	return input[m-1][n-1]
}

package main

import (
	"fmt"
)

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	maxKarakter := h[word[0]-97]

	for i := 1; i < len(word); i++ {
		if maxKarakter < h[word[i] - 97] {
			maxKarakter = h[word[i]-97]
		}
	}
	result := len(word) * int(maxKarakter)
	return int32(result)
}

func main()  {
	h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	word := "abc"
	num := designerPdfViewer(h, word)
    fmt.Println(num)
}
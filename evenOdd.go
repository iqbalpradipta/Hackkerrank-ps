package main

import "fmt"

func EvenOrOdd(number int) string {
  // your code here
  str1 := "Odd"
  str2 := "Even"
  var str string
  for i:= number ; i <= number; i++ {
    if i % 2 != 0 {
      str = str1
    } else {
      str = str2
    }
  }

  return str
}

func main()  {
	number := -1
	fmt.Println(EvenOrOdd(number))
}
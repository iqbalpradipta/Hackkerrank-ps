package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) string {
	strSeparate := strings.Fields(str)
	for i, v := range strSeparate {
		strSeparate[i] = strings.ToUpper(string(v[0])) + v[1:]
	}
	return strings.Join(strSeparate, " ")
}

func main()  {
	str := "How can mirrors be real if our eyes aren't real"
	fmt.Println(ToJadenCase(str))
}
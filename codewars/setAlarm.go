package main

import "fmt"

func SetAlarm(employed, vacation bool) bool {
	if employed != true || employed == true && vacation == true{
		return false
	}else {
		return true
	}
}

func  main()  {
	employed := true
	vacation := true
	fmt.Println(SetAlarm(employed, vacation))
}
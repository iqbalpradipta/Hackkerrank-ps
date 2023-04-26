package main

// import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    // Write your code here
	if v1 <= v2 {
        return "NO"
    } else {
        if (x2 - x1) % (v1 - v2) == 0{
            return "YES"
        } else {
            return "NO"
        }
    }
}

// func main()  {
//     fmt.Println(kangaroo(0,3,4,2)) //YES
//     fmt.Println(kangaroo(0,2,5,3)) //
    
// }


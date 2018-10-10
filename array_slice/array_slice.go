package main

import (
	"fmt"
)

// func main() {
// 	var a [5]int
// 	fmt.Println(a)
// 	a[2] = 100
// 	fmt.Println(a)
// 	// a[5] = 1000
// 	// fmt.Println(a)
// }

func main() {
	i := make([]int, 3)
	fmt.Printf("Capacity of i is %d & length is %d\n", cap(i), len(i))
	i[1] = 100
	fmt.Printf("Capacity of i is %d & length is %d\n", cap(i), len(i))
	fmt.Println(i)
	i = append(i, 1000, -999, 4444)
	fmt.Println(i)
	fmt.Printf("Capacity of i is %d & length is %d\n", cap(i), len(i))

}

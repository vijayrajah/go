package main

import (
	"fmt"
)

func main() {
	num := []int{10, 21, 32, 43, 54, 65}

	for a, b := range num {
		// fmt.Println("A is ", a, "B is ", b)
		fmt.Println("num[a] is ", num[a], " and b is ", b)
	}

	strin := make(map[string]string)

	strin["ONE"] = "1"
	strin["Two"] = "2"

	for a, b := range strin {
		fmt.Println("strin[a] is ", strin[a], " and b is ", b)
	}
}

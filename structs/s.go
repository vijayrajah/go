package main

import (
	"fmt"
)

type mystruct struct {
	sid string
	num int
}

func main() {

	a := mystruct{sid: "1-2-3-5", num: 20}
	b := mystruct{sid: "A-B-C-5", num: 100}

	fmt.Println("a is ", a, "b is ", b)

	c := b.num
	fmt.Println("C is ", c)

	sarr := make([]mystruct, 10)
	sarr[0].num = 1
	sarr[0].sid = "ABCD"
	fmt.Println("sarr is ", sarr)
}

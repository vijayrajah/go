package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)

	m["A"] = "One"
	m["b"] = "two"

	a, prs := m["A"]
	fmt.Println("A is ", a, "prs is ", prs)

	a, prs = m["C"]
	fmt.Println("A is ", a, "prs is ", prs)

	fmt.Println(m)
	fmt.Println(m["a"])
	fmt.Println(m["A"])
	fmt.Println(m["c"])
}

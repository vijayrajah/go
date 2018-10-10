package main

import (
	"fmt"
	"strconv"
)

//GLOBAL VARIABLE
var count int

type s1 struct {
	ident   int
	name    string
	age     int
	address string
}

func printAge(s1 s1) {
	fmt.Println("Age is ", s1.age)
}

func main() {

	count++
	persona := s1{ident: count, name: "Vijay", age: 35, address: "Chennai, India"}
	printAge(persona)

	personarr := make([]s1, 0)
	for i := 0; i < 10; i++ {
		count++
		personarr = append(personarr, s1{ident: count, name: "Vijay" + strconv.Itoa(i), age: 35 + i, address: "Chennai, India"})
	}

	fmt.Println(personarr)

}

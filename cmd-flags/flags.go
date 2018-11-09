package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	strInp = kingpin.Flag("string", "Input a string").Short('s').Required().String()
	iInp   = kingpin.Flag("integ", "Give an integer").Required().Int()
	stropt = kingpin.Flag("opti", "optional input").Short('o').Default("abcdef").String()
	boOpt  = kingpin.Flag("bo", "Boolean parameter").Short('b').Default("false").Bool()
)

func main() {
	kingpin.Parse()
	fmt.Println("strinp is ", *strInp, " And iInp is ", *iInp, " and optional arg is ", *stropt, " and boolopt is ", *boOpt)
}

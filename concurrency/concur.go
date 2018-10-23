package main

import (
	"fmt"
	"time"
)

func funca() {

	fmt.Println("in funca :: Sleeping for ...  100 Milli secs")
	time.Sleep(time.Millisecond * 100)
}

func funcb() {

	fmt.Println("in funcb :: Sleeping for ... * 1000 Milli secs")
	time.Sleep(time.Millisecond * 1000)
}

func funcc(inp <-chan string) {
	fmt.Println("the channel input is ", <-inp)
}

func main() {

	funccInp := make(chan string)
	go func() {
		funccInp <- "from main"
	}()

	go funcc(funccInp)
	go funcb()
	go funca()

}

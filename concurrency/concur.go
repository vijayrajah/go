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
	//funccInp <- "from main" //This will block and you will get "fatal error: all goroutines are asleep - deadlock!"
	//This is because the channel defined is a syncronous channel, and hence, it will be blocked untill it read from
	//Hence, the main thread is blocked and go detecs it and throws the error

	// we can however, write to the channel in async way like....
	go func() {
		funccInp <- "from main"
	}()
	//This does not block the main thread, as a seperate go routine is used..

	go funcc(funccInp)
	go funcb()
	go funca()

}

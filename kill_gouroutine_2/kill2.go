package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	t := 8

	kill := make(chan int)

	var wg sync.WaitGroup

	//start the go routine
	wg.Add(1)
	go someSleepfunc(kill, &wg)

	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("killing.. ", time.Now().String())
	close(kill)

	//now wait
	wg.Wait()

}

func someSleepfunc(kill chan int, wg *sync.WaitGroup) {

	var i int

	for {

		select {
		case <-kill:
			wg.Done() //Need to decrement the counter.. otherwise we will get a deadlock
			return
		case <-time.After(1 * time.Second):
			fmt.Println("The current time is ", time.Now().String())
			i++
			if i > 2 {
				wg.Done()
				return
			}
		}
	}

}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	//generate random sequence
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	t := r1.Intn(10)

	fmt.Println("Wait time is ", t)

	kill := make(chan int)

	var wg sync.WaitGroup

	//start the go routine
	wg.Add(1)
	go someSleepfunc(kill, &wg)

	time.Sleep(time.Duration(t) * time.Second)
	close(kill)

	//now wait
	wg.Wait()

}

func someSleepfunc(kill chan int, wg *sync.WaitGroup) {

	for {

		select {
		case <-kill:
			wg.Done() //Need to decrement the counter.. otherwise we will get a deadlock
			return
		case <-time.After(1 * time.Second):
			fmt.Println("The current time is ", time.Now().String())
		}
	}

}

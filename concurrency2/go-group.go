package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	wg1 sync.WaitGroup
	wg2 sync.WaitGroup
)

func fa(delay int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Sleeping for ", strconv.Itoa(delay), " Secs... in fa")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("slept for ", strconv.Itoa(delay), " Secs... in fa")
}

func fb(delay int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Sleeping for ", strconv.Itoa(delay), " Secs... in fb")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("slept for ", strconv.Itoa(delay), " Secs... in fb")
}

func fc(delay int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Sleeping for ", strconv.Itoa(delay), " Secs... in fc")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("slept for ", strconv.Itoa(delay), " Secs... in fc")
}

func fd(delay int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Sleeping for ", strconv.Itoa(delay), " Secs... in fd")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("slept for ", strconv.Itoa(delay), " Secs... in fd")

}

func fe(delay int) {

	fmt.Println("Sleeping for ", strconv.Itoa(delay), " Secs... in fe")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("slept for ", strconv.Itoa(delay), " Secs... in fe")
}

func main() {
	//rand seed
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	done := make(chan bool, 2)

	wg1.Add(2)
	go fa(rand.Intn(10), &wg1)
	go fb(rand.Intn(10), &wg1)

	wg2.Add(2)
	go fc(rand.Intn(10), &wg2)
	go fd(rand.Intn(10), &wg2)

	go func(wg *sync.WaitGroup, done chan bool) {
		wg.Wait()
		fmt.Println("group 1 is done... Calling fe")
		go fe(rand.Intn(10))

		done <- true

	}(&wg1, done)

	go func(wg *sync.WaitGroup, done chan bool) {
		wg.Wait()
		fmt.Println("group 2 is done.")

		done <- true

	}(&wg2, done)

	<-done
}

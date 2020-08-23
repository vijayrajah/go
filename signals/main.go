package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	sig := make(chan os.Signal, 1)

	//create a context so we can cancel it later
	ctx, cancell := context.WithCancel(context.Background())

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	doSomethingInital()

	//start a goroutine and wait for signal
	go func() {
		//wait for signals
		<-sig

		//once we get the signal.. Call cancell()
		cancell()
	}()

	//start func 1 on a timer
	wg.Add(1)
	ticker1 := time.NewTicker(3 * time.Second)
	go func(wg *sync.WaitGroup, c1 context.Context) {
		defer wg.Done()
		for {
			select {
			case <-c1.Done():
				{
					fmt.Println("Func1 done")
					return
				}
			case <-ticker1.C:
				{
					someWorkerFunc1()
				}

			}
		}
	}(&wg, ctx)

	//start func 2 on a timer
	wg.Add(1)
	ticker2 := time.NewTicker(2 * time.Second)
	go func(wg *sync.WaitGroup, c2 context.Context) {
		defer wg.Done()
		for {
			select {
			case <-c2.Done():
				{
					fmt.Println("Func2 done")
					return
				}
			case <-ticker2.C:
				{
					someWorkerFunc2()
				}

			}
		}
	}(&wg, ctx)

	//now wait
	wg.Wait()
	fmt.Println("All done")

}

func doSomethingInital() {
	fmt.Println("Some initial stuff")
}

func someWorkerFunc1() {
	time.Sleep(1 * time.Second)
	fmt.Println("Func1: The current time is", time.Now().String())
}

func someWorkerFunc2() {
	time.Sleep(2 * time.Second)
	fmt.Println("Func2: The current time is", time.Now().String())
}

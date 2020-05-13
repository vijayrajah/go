package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//start a pool of workers to do some unit of work

//some struct to hold data
type jobResult struct {
	start     time.Time
	end       time.Time
	sleepTime int
}

//a mock function to do 'actual' work
func actualJob() int {
	//generate random sequence
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	t := r1.Intn(5)
	time.Sleep(time.Duration(t) * time.Second)
	return t

}

//worker function...
func worker(id int, jobs <-chan int, wg *sync.WaitGroup, r []jobResult) {

	defer wg.Done()

	for j := range jobs {

		startime := time.Now()

		fmt.Println("worker", id, "started  job", j, "My index is ", j)
		t := actualJob()
		r[j] = jobResult{ //j is guarnteed to be unique since it is the job index
			start:     startime,
			sleepTime: t,
			end:       time.Now(),
		}
		fmt.Println("worker", id, "finished job", j)

	}
}

func main() {

	//set Basic vars
	numJobs := 5
	maxworkers := 3
	var wg sync.WaitGroup
	results := make([]jobResult, numJobs)

	jobs := make(chan int, numJobs)

	//start 'maxworkers' workers asyncronously
	// as we send some data to channel, they will start
	// here they will start and be blocked...
	for w := 0; w < maxworkers; w++ {
		wg.Add(1) //goroutine counting
		go worker(w, jobs, &wg, results)
	}

	//send actual job to workers....
	for j := 0; j < numJobs; j++ {
		jobs <- j
	}

	//close channel.. all the work is completed
	close(jobs)

	// wait for all the goroutines to complete
	wg.Wait()

	//now print results
	for i, a := range results {
		fmt.Println("Job: ", i, " Start : ", a.start.String(), " End: ", a.end.String(), " Sleptfor: ", a.sleepTime)
	}
}

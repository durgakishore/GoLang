package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Job struct
type Job struct {
	id       int
	randomno int
}

//Result struct
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {

	sum := 0
	for number != 0 {
		digits := number % 10
		sum += digits * digits
		number /= 10
	}
	//	time.Sleep(2 * time.Second)
	return sum
}

func allocate(nofJobs int) {
	for i := 0; i < nofJobs; i++ {
		randNumber := rand.Intn(999)
		job := Job{i, randNumber}
		jobs <- job
	}
	close(jobs)
}

func createWorkerPool(nofWorkers int) {

	var wg sync.WaitGroup

	for i := 0; i < nofWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup) {

	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func printResults(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {

	starttime := time.Now()

	numberofJobs := 100
	go allocate(numberofJobs)
	done := make(chan bool)
	go printResults(done)
	numberofWorkers := 10
	createWorkerPool(numberofWorkers)
	<-done
	endtime := time.Now()
	diff := endtime.Sub(starttime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

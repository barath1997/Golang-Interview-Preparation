package main

import "sync"

type Job struct {
	JobID    int
	JobValue int
}

func worker(jobs chan Job, arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for ip := range jobs {
		arr[ip.JobID] = ip.JobValue * 2
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	jobs := make(chan Job)
	//results := make(chan int)
	wg := &sync.WaitGroup{}

	minWorkers := 3
	workerCount := min(minWorkers, len(arr)/3)
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(jobs, arr, wg)
	}

	go func() {
		for idx, v := range arr {
			jobs <- Job{JobID: idx, JobValue: v}
		}
		close(jobs)
	}()

	wg.Wait()

	for _, val := range arr {
		println(val)
	}

}

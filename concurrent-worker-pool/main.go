package main

import "sync"

func worker(jobs chan int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for ip := range jobs {
		results <- ip * 2
	}
}

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	jobs := make(chan int)
	results := make(chan int)
	wg := &sync.WaitGroup{}

	N := 3
	for i := 1; i <= N; i++ {
		wg.Add(1)
		go worker(jobs, results, wg)
	}

	go func() {
		for _, v := range arr {
			jobs <- v
		}
		close(jobs)
	}()

	for res := range results {
		println(res)
	}

	wg.Wait()
	close(results)
}
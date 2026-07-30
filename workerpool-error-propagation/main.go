package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func worker(ip <-chan int, op chan<- int, err chan error, ctx context.Context, value int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val, ok := <-ip:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond * 500)
			op <- val * value
		case <-ctx.Done():
			w := fmt.Sprintf("worker %d stopped, time completed!!", value)
			err <- errors.New(w)
			return
		}
	}
}

func consumer(op chan int, err chan error, wg *sync.WaitGroup, errArr *[]error, resultArr *[]int) {
	defer wg.Done()
	for {
		if op == nil && err == nil {
			return
		}
		select {
		case output, ok := <-op:
			if !ok {
				op = nil
				continue
			}
			*resultArr = append(*resultArr, output)
		case er, ok := <-err:
			if !ok {
				err = nil
				continue
			}
			*errArr = append(*errArr, er)
		}
	}
}

func main() {
	ip := make(chan int)
	op := make(chan int)
	errCh := make(chan error)
	resultArr := make([]int, 0)
	errArr := make([]error, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	wgWorkers := &sync.WaitGroup{}
	wgConsumer := &sync.WaitGroup{}

	// workers
	for i := 1; i <= 3; i++ {
		wgWorkers.Add(1)
		go func(value int) {
			worker(ip, op, errCh, ctx, value, wgWorkers)
		}(i)
	}

	// consumer
	wgConsumer.Add(1)
	go consumer(op, errCh, wgConsumer, &errArr, &resultArr)

	// jobs
	for i := 1; i <= 5; i++ {
		ip <- i
	}
	close(ip)

	// wait only for workers, then close channels
	wgWorkers.Wait()
	close(op)
	close(errCh)

	// now wait for consumer to drain & finish
	wgConsumer.Wait()

	// safe to read results
	fmt.Println("errors:", errArr)
	fmt.Println("results:", resultArr)
}

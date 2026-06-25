package main

import (
	"context"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for i := range 5 {
			select {
			case <-ctx.Done():
				return
			case ch1 <- i:
				time.Sleep(time.Millisecond * 1000)
			}
		}
		close(ch1)
	}()

	go func() {
		for i := range 5 {
			select {
			case <-ctx.Done():
				return
			case ch1 <- i:
				time.Sleep(time.Millisecond * 500)
			}

		}
		close(ch2)
	}()

	valueCount := 0

	for valueCount < 5 {
		select {
		case val1 := <-ch1:
			valueCount += 1
			println("ch1 : ", val1)
		case val2 := <-ch2:
			valueCount += 1
			println("ch2 : ", val2)
		}
	}

	// loop ends
	cancel()

}

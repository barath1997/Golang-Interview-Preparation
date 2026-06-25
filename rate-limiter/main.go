package main

import (
	"time"
)

func rateLimit(ch <-chan int) {

	// let the 'x' miilliseconds be 500ms

	for value := range ch {
		println(value * 2)
		time.Sleep(time.Millisecond * 500)
	}

}

func main() {

	ch := make(chan int)
	go rateLimit(ch)

	for i := range 6 {
		ch <- i
	}
	close(ch)

}

package main

import (
	"time"
)

func rateLimit(ch <-chan int, ticker *time.Ticker) {

	defer ticker.Stop()
	for value := range ch {
		println(value * 2)
		<-ticker.C

	}

}

func main() {

	ch := make(chan int)
	ticker := time.NewTicker(time.Millisecond * 500)
	go rateLimit(ch, ticker)

	for i := range 6 {
		ch <- i
	}
	close(ch)

}

package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

func work(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	select {
	case <-time.After(time.Second * 3):
		println("all is well")
		return nil
	case <-ctx.Done():
		println("error")
		return errors.New("timeout")
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	wg.Add(1)
	go work(ctx, wg)
	wg.Wait()

}

package main

import "sync"

type counter struct {
	mu    sync.Mutex
	value int
}

func (c *counter) Inc() {
	c.mu.Lock()
	c.value += 1
	c.mu.Unlock()
}

func (c *counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	c := new(counter)
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	println(c.Value())
}

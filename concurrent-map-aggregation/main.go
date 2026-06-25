package main

import "sync"

type count struct {
	mu      sync.Mutex
	counter map[string]int
}

func main() {
	c := count{counter: make(map[string]int)}
	input := []string{"hell", "heaven", "new", "baby", "heaven", "baby"}
	wg := &sync.WaitGroup{}

	for _, value := range input {
		wg.Add(1)
		go func(value string) {
			defer wg.Done()
			if _, ok := c.counter[value]; ok {
				c.mu.Lock()
				defer c.mu.Unlock()
				c.counter[value]++
			} else {
				c.mu.Lock()
				defer c.mu.Unlock()
				c.counter[value] = 1

			}

		}(value)
	}
	wg.Wait()

	for key, value := range c.counter {
		println(key, value)
	}

}

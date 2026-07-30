package main

import (
	"context"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*1)
	defer cancel()

}

func FetchAll(ctx context.Context, urls []string, maxConcurrency, maxRetries int) (map[string][]byte, error) {
	deadline, _ := ctx.Deadline()

	client := &http.Client{Timeout: time.Until(deadline)}

	for i := 0; i < maxConcurrency; i++ {

	}

}

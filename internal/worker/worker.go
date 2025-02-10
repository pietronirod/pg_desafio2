package worker

import (
	"net/http"
	"sync"
	"time"

	"load-tester/internal/types"
)

func Worker(id int, url string, requests int, results chan<- types.Result, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{}

	for i := 0; i < requests; i++ {
		start := time.Now()
		resp, err := client.Get(url)
		duration := time.Since(start)

		if err != nil {
			results <- types.Result{StatusCode: 0, Duration: duration}
			continue
		}

		results <- types.Result{StatusCode: resp.StatusCode, Duration: duration}
		resp.Body.Close()
	}
}

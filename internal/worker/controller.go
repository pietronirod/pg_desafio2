package worker

import (
	"load-tester/internal/report"
	"load-tester/internal/types"
	"sync"
	"time"
)

func RunLoadTest(url string, totalRequests int, concurrency int) {
	var wg sync.WaitGroup
	results := make(chan types.Result, totalRequests)

	requestsPerWorker := totalRequests / concurrency
	remainder := totalRequests % concurrency

	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)

		numRequests := requestsPerWorker
		if i < remainder {
			numRequests++
		}

		go Worker(i, url, numRequests, results, &wg)
	}

	wg.Wait()
	close(results)

	endTime := time.Since(startTime)

	report.ProcessResults(results, endTime)
}

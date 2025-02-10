package worker

import (
	"fmt"
	"sync"
	"time"
)

func RunLoadTest(url string, totalRequests int, concurrency int) {
	var wg sync.WaitGroup
	results := make(chan Result, totalRequests)

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

	processResults(results, endTime)
}

func processResults(results chan Result, duration time.Duration) {
	statusCount := make(map[int]int)
	totalRequests := 0

	for result := range results {
		statusCount[result.StatusCode]++
		totalRequests++
	}

	fmt.Println("\n--- Relatório Final ---")
	fmt.Printf("Tempo total: %v\n", duration)
	fmt.Printf("Total de requests: %d\n", totalRequests)

	for status, count := range statusCount {
		fmt.Printf("HTTP %d: %d vezes\n", status, count)
	}
}

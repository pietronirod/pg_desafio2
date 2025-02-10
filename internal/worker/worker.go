package worker

import (
	"log"
	"net/http"
	"sync"
	"time"

	"load-tester/internal/types"
)

func Worker(id int, url string, requests int, results chan<- types.Result, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for i := 0; i < requests; i++ {
		var statusCode int
		start := time.Now()

		for retry := 0; retry < 3; retry++ {
			resp, err := client.Get(url)
			duration := time.Since(start)

			if err != nil {
				log.Printf("[Worker %d] Erro na requisicão (tentativa %d): %v\n", id, retry+1, err)
				time.Sleep(500 * time.Millisecond)
				continue
			}

			statusCode = resp.StatusCode
			resp.Body.Close()
			results <- types.Result{StatusCode: resp.StatusCode, Duration: duration}
			break
		}

		if statusCode == 0 {
			results <- types.Result{StatusCode: -1, Duration: 5 * time.Second}
		}
	}
}

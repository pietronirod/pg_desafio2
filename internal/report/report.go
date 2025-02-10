package report

import (
	"fmt"
	"time"

	"load-tester/internal/types"
)

func ProcessResults(results chan types.Result, duration time.Duration) {
	statusCount := make(map[int]int)
	var totalRequests int
	var totalDuration time.Duration
	var timeouts int

	for result := range results {
		if result.StatusCode == -1 {
			timeouts++
			continue
		}

		statusCount[result.StatusCode]++
		totalRequests++
		totalDuration += result.Duration
	}

	var avgDuration time.Duration
	if totalRequests > 0 {
		avgDuration = totalDuration / time.Duration(totalRequests)
	}

	fmt.Println("\n--- Relatório Final ---")
	fmt.Printf("Tempo total: %v\n", duration)
	fmt.Printf("Total de requests: %d\n", totalRequests)
	fmt.Printf("Tempo médio de resposta: %v\n", avgDuration)
	fmt.Printf("Timeouts: %d\n", timeouts)

	fmt.Println("\nDistribuicão de Status HTTP:")
	for status, count := range statusCount {
		fmt.Printf(" - HTTP %d: %d vezes\n", status, count)
	}

	fmt.Println("\nTeste concluído!")
}

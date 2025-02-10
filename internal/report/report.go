package report

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"load-tester/internal/types"
)

func ProcessResults(results []types.Result, duration time.Duration, outputFormat string) {
	statusCount := make(map[int]int)
	var totalRequests int
	var totalDuration time.Duration
	var timeouts int

	for _, result := range results {
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

	switch outputFormat {
	case "json":
		saveJSON(results)
	case "csv":
		saveCSV(results)
	default:
		printReport(duration, totalRequests, avgDuration, timeouts, statusCount)
	}
}

func printReport(duration time.Duration, totalRequests int, avgDuration time.Duration, timeouts int, statusCount map[int]int) {
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

func saveJSON(results []types.Result) {
	file, _ := os.Create("report.json")
	defer file.Close()

	jsonData, _ := json.MarshalIndent(results, "", " ")
	file.Write(jsonData)

	fmt.Println("Relatório salvo como report.json")
}

func saveCSV(results []types.Result) {
	file, _ := os.Create("report.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Status Code", "Duration (ms)"})

	for _, result := range results {
		writer.Write([]string{fmt.Sprintf("%d", result.StatusCode), fmt.Sprintf("%d", result.Duration.Microseconds())})
	}

	fmt.Println("Relatório salvo como report.csv")
}

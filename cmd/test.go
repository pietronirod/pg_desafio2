/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"load-tester/internal/report"
	"load-tester/internal/worker"

	"github.com/spf13/cobra"
)

var url string
var requests int
var concurrency int
var outputFormat string

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Executa um teste de carga",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" || requests <= 0 || concurrency <= 0 {
			fmt.Println("Parâmetros inválidos. Use --url, --requests e --concurrency corretamente.")
			return
		}

		fmt.Println("Iniciando teste de carga...")
		fmt.Printf("URL: %s\n", url)
		fmt.Printf("Total de Requests: %d\n", requests)
		fmt.Printf("Concorrência: %d\n", concurrency)

		results, duration := worker.RunLoadTest(url, requests, concurrency)

		report.ProcessResults(results, duration, outputFormat)
	},
}

func init() {
	testCmd.Flags().StringVar(&url, "url", "", "URL do serviço")
	testCmd.Flags().IntVar(&requests, "requests", 10, "Número total de requests")
	testCmd.Flags().IntVar(&concurrency, "concurrency", 2, "Número de chamadas simultâneas")
	testCmd.Flags().StringVar(&outputFormat, "output", "console", "Formato de saída: console, json ou csv")
	rootCmd.AddCommand(testCmd)
}

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Saad7890-web/go-log-analyzer/internal/analyzer"
	"github.com/Saad7890-web/go-log-analyzer/internal/parser"
)

func main() {
	file, err := os.Open("testdata/sample.log")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var entries []parser.LogEntry

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		entry, err :=parser.ParseLine(line)
		if err == nil {
			entries = append(entries, entry)
		}
	}

	report := analyzer.Analyze(entries)
	fmt.Println("===== Log Report =====")
	fmt.Println("Total Requests:", report.TotalRequests)
	fmt.Println("Total Errors:", report.TotalErrors)
	fmt.Println("Average Response Time:", report.AvgResponse, "ms")
	fmt.Println("Top IP:", report.TopIP)
	fmt.Println("Top Endpoint:", report.TopEndpoint)


}
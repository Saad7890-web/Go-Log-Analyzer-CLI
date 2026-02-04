package analyzer

import (
	"testing"

	"github.com/Saad7890-web/go-log-analyzer/internal/parser"
)

func TestAnalyze(t *testing.T) {
	entries := []parser.LogEntry{
		{IP: "1.1.1.1", Endpoint: "/a", StatusCode: 200, ResponseTime: 100},
		{IP: "1.1.1.1", Endpoint: "/a", StatusCode: 500, ResponseTime: 200},
		{IP: "2.2.2.2", Endpoint: "/b", StatusCode: 200, ResponseTime: 300},
	}

	report := Analyze(entries)

	if report.TotalRequests != 3 {
		t.Errorf("TotalRequests got %d, want 3", report.TotalRequests)
	}

	if report.TotalErrors != 1 {
		t.Errorf("TotalErrors got %d, want 1", report.TotalErrors)
	}

	if report.TopIP != "1.1.1.1" {
		t.Errorf("TopIP got %s, want 1.1.1.1", report.TopIP)
	}

	if report.TopEndpoint != "/a" {
		t.Errorf("TopEndpoint got %s, want /a", report.TopEndpoint)
	}
}

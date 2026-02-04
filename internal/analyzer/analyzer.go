package analyzer

import "github.com/Saad7890-web/go-log-analyzer/internal/parser"
	

type Report struct {
	TotalRequests int
	TotalErrors int
	AvgResponse float64
	TopIP string
	TopEndpoint string
}


func Analyze(entries []parser.LogEntry) Report {
	ipCount := make(map[string]int)
	endpointCount := make(map[string]int)

	var totalResp int
	var errors int

	for _,e := range entries {
		ipCount[e.IP]++
		endpointCount[e.Endpoint]++
		totalResp += e.ResponseTime

		if e.StatusCode >= 500 {
			errors++;
		}
	}

	return Report{
		TotalRequests: len(entries),
		TotalErrors: errors,
		AvgResponse: float64(totalResp)/float64(len(entries)),
		TopIP: maxKey(ipCount),
		TopEndpoint: maxKey(endpointCount),
	}
}

func maxKey(m map[string]int) string {
	var maxKey string
	var maxVal int

	for k,v := range m {
		if v > maxVal {
			maxVal = v
			maxKey = k
		}
	}

	return maxKey
}

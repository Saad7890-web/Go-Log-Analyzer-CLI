package parser

import (
	"strconv"
	"strings"
)

type LogEntry struct {
	IP string
	Method string
	Endpoint string
	StatusCode int
	ResponseTime int 
}

func ParseLine(line string)(LogEntry, error){
	parts := strings.Fields(line)

	status, err := strconv.Atoi(parts[3])
	if err != nil {
		return LogEntry{}, err
	}

	time,err:= strconv.Atoi(parts[4])
	if err != nil {
		return LogEntry{}, err
	}

	return LogEntry{
		IP: parts[0],
		Method: parts[1],
		Endpoint: parts[2],
		StatusCode: status,
		ResponseTime: time,
	}, nil
}
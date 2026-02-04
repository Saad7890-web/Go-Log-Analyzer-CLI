package parser

import "testing"

func TestParseLine(t *testing.T){
	tests := []struct {
		line string
		want LogEntry
	}{
		{
			line: "192.168.0.1 GET /api/users 200 120",
			want: LogEntry{
				IP: "192.168.0.1",
				Method: "GET",
				Endpoint: "/api/users",
				StatusCode: 200,
				ResponseTime: 120,
			},
		},
		{
			line: "10.0.0.1 POST /api/login 500 300",
			want: LogEntry{
				IP:           "10.0.0.1",
				Method:       "POST",
				Endpoint:     "/api/login",
				StatusCode:   500,
				ResponseTime: 300,
			},
		},
	}

	for _, tt := range tests {
		got, err := ParseLine(tt.line)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != tt.want {
			t.Errorf("got %+v, want %+v", got, tt.want)
		}
	}
}
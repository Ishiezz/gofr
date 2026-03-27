package service

import (
	"fmt"
	"io"
	"time"
)

type Logger interface {
	Log(args ...any)
}

type Log struct {
	Timestamp     time.Time `json:"timestamp"`
	ResponseTime  any       `json:"latency"`
	CorrelationID string    `json:"correlationId"`
	ResponseCode  int       `json:"responseCode"`
	HTTPMethod    string    `json:"httpMethod"`
	URI           string    `json:"uri"`
}

func (l *Log) PrettyPrint(writer io.Writer) {
	var responseTime int64

	switch v := l.ResponseTime.(type) {
	case int64:
		responseTime = v
	case int:
		responseTime = int64(v)
	case string:
		fmt.Sscanf(v, "%d", &responseTime)
	}

	fmt.Fprintf(writer, "\u001B[38;5;8m%s \u001B[38;5;%dm%-6d\u001B[0m %8d\u001B[38;5;8mµs\u001B[0m %s %s \n",
		l.CorrelationID, colorForStatusCode(l.ResponseCode),
		l.ResponseCode, responseTime, l.HTTPMethod, l.URI)
}

type ErrorLog struct {
	*Log
	ErrorMessage string `json:"errorMessage"`
}

func (el *ErrorLog) PrettyPrint(writer io.Writer) {
	var responseTime int64

	switch v := el.ResponseTime.(type) {
	case int64:
		responseTime = v
	case int:
		responseTime = int64(v)
	case string:
		fmt.Sscanf(v, "%d", &responseTime)
	}

	fmt.Fprintf(writer, "\u001B[38;5;8m%s \u001B[38;5;%dm%-6d\u001B[0m %8d\u001B[38;5;8mµs\u001B[0m %s %s \n",
		el.CorrelationID, colorForStatusCode(el.ResponseCode),
		el.ResponseCode, responseTime, el.HTTPMethod, el.URI)
}

func colorForStatusCode(status int) int {
	const (
		blue   = 34
		red    = 202
		yellow = 220
	)

	switch {
	case status >= 200 && status < 300:
		return blue
	case status >= 400 && status < 500:
		return yellow
	case status >= 500 && status < 600:
		return red
	}

	return 0
}

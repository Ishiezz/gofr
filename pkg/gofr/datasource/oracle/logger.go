package oracle

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Logger interface {
	Debugf(pattern string, args ...any)
	Debug(args ...any)
	Logf(pattern string, args ...any)
	Errorf(pattern string, args ...any)
}

type Log struct {
	Type     string `json:"type"`
	Query    string `json:"query"`
	Duration any    `json:"duration"`
	Args     []any  `json:"args,omitempty"`
}

func (l *Log) PrettyPrint(writer io.Writer) {
	var duration int64

	switch v := l.Duration.(type) {
	case int64:
		duration = v
	case int:
		duration = int64(v)
	case string:
		fmt.Sscanf(v, "%d", &duration)
	}

	fmt.Fprintf(writer, "%-10s ORACLE %8dµs %s\n", l.Type, duration, clean(l.Query))
}

func clean(query string) string {
	query = regexp.MustCompile(`\s+`).ReplaceAllString(query, " ")
	return strings.TrimSpace(query)
}

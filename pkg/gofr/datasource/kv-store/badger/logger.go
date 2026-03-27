package badger

import (
	"fmt"
	"io"
)

type Logger interface {
	Debug(args ...any)
	Debugf(pattern string, args ...any)
	Info(args ...any)
	Infof(pattern string, args ...any)
	Error(args ...any)
	Errorf(pattern string, args ...any)
}

type Log struct {
	Type     string `json:"type"`
	Duration any    `json:"duration"`
	Key      string `json:"key"`
	Value    string `json:"value,omitempty"`
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

	fmt.Fprintf(writer, "\u001B[38;5;8m%-32s \u001B[38;5;162m%-6s\u001B[0m %8d\u001B[38;5;8mµs\u001B[0m %s \n",
		l.Type, "BADGR", duration, l.Key+" "+l.Value)
}

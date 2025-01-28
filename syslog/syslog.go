package syslog

import (
	"fmt"
	"log/syslog"
)

type Logger struct {
	writer *syslog.Writer
}

func NewLogger(host string, port int) (*Logger, error) {
	writer, err := syslog.Dial("udp", fmt.Sprintf("%s:%d", host, port), syslog.LOG_INFO, "syslog-gen")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to syslog host: %v", err)
	}
	return &Logger{writer: writer}, nil
}

func (l *Logger) Close() {
	l.writer.Close()
}

func (l *Logger) Send(message string, level syslog.Priority) {
	switch level {
	case syslog.LOG_INFO:
		l.writer.Info(message)
	case syslog.LOG_ERR:
		l.writer.Err(message)
	case syslog.LOG_WARNING:
		l.writer.Warning(message)
	case syslog.LOG_CRIT:
		l.writer.Crit(message)
	}
}

package syslog

import (
	"fmt"
	"net"

	"github.com/eynsfordcq/go-syslog-simulator/config"
)

type Logger struct {
	conn   net.Conn
	config *config.Config
}

func NewLogger(cfg *config.Config) (*Logger, error) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to syslog host: %v", err)
	}
	return &Logger{
		conn:   conn,
		config: cfg,
	}, nil
}

func (l *Logger) Close() {
	l.conn.Close()
}

func (l *Logger) Send(message string) {
	l.conn.Write([]byte(message))
}

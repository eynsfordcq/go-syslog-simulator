package config

import (
	"encoding/json"
	"fmt"
	"log/syslog"
	"os"
)

type Config struct {
	Hostname           string   `json:"hostname"`
	DomainName         string   `json:"domainName"`
	Tags               []string `json:"tags"`
	SyslogLevels       []string `json:"syslogLevels"`
	SampleLogs         []string `json:"sampleLogs"`
	RequestsPerSecond  int
	Host               string
	Port               int
	ParsedSyslogLevels []syslog.Priority
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	err = cfg.ParseSyslogLevels()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) ParseSyslogLevels() error {
	levels := make([]syslog.Priority, 0, len(c.SyslogLevels))
	for _, level := range c.SyslogLevels {
		switch level {
		case "LOG_INFO":
			levels = append(levels, syslog.LOG_INFO)
		case "LOG_ERR":
			levels = append(levels, syslog.LOG_ERR)
		case "LOG_WARNING":
			levels = append(levels, syslog.LOG_WARNING)
		case "LOG_CRIT":
			levels = append(levels, syslog.LOG_CRIT)
		default:
			return fmt.Errorf("unknown syslog level: %s", level)
		}
	}
	c.ParsedSyslogLevels = levels
	return nil
}

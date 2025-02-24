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
	Facility           string   `json:"facility"`
	RequestsPerSecond  int
	Host               string
	Port               int
	ParsedSyslogLevels []syslog.Priority
	ParsedFacility     syslog.Priority
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

	err = cfg.ParseFacility()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) ParseSyslogLevels() error {
	levels := make([]syslog.Priority, 0, len(c.SyslogLevels))
	for _, level := range c.SyslogLevels {
		switch level {
		case "LOG_EMERG":
			levels = append(levels, syslog.LOG_EMERG)
		case "LOG_ALERT":
			levels = append(levels, syslog.LOG_ALERT)
		case "LOG_CRIT":
			levels = append(levels, syslog.LOG_CRIT)
		case "LOG_ERR":
			levels = append(levels, syslog.LOG_ERR)
		case "LOG_WARNING":
			levels = append(levels, syslog.LOG_WARNING)
		case "LOG_NOTICE":
			levels = append(levels, syslog.LOG_NOTICE)
		case "LOG_INFO":
			levels = append(levels, syslog.LOG_INFO)
		case "LOG_DEBUG":
			levels = append(levels, syslog.LOG_DEBUG)
		default:
			return fmt.Errorf("unknown syslog level: %s", level)
		}
	}
	c.ParsedSyslogLevels = levels
	return nil
}

func (c *Config) ParseFacility() error {
	switch c.Facility {
	case "LOG_KERN":
		c.ParsedFacility = syslog.LOG_KERN
	case "LOG_USER":
		c.ParsedFacility = syslog.LOG_USER
	case "LOG_MAIL":
		c.ParsedFacility = syslog.LOG_MAIL
	case "LOG_DAEMON":
		c.ParsedFacility = syslog.LOG_DAEMON
	case "LOG_AUTH":
		c.ParsedFacility = syslog.LOG_AUTH
	case "LOG_SYSLOG":
		c.ParsedFacility = syslog.LOG_SYSLOG
	case "LOG_LPR":
		c.ParsedFacility = syslog.LOG_LPR
	case "LOG_NEWS":
		c.ParsedFacility = syslog.LOG_NEWS
	case "LOG_UUCP":
		c.ParsedFacility = syslog.LOG_UUCP
	case "LOG_CRON":
		c.ParsedFacility = syslog.LOG_CRON
	case "LOG_AUTHPRIV":
		c.ParsedFacility = syslog.LOG_AUTHPRIV
	case "LOG_FTP":
		c.ParsedFacility = syslog.LOG_FTP
	default:
		return fmt.Errorf("unknown syslog facility: %s", c.Facility)
	}
	return nil
}

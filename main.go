package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eynsfordcq/go-syslog-simulator/config"
	"github.com/eynsfordcq/go-syslog-simulator/syslog"
	"github.com/eynsfordcq/go-syslog-simulator/utils"
)

func sendSyslogs(logger *syslog.Logger, cfg *config.Config) {
	interval := time.Second / time.Duration(cfg.RequestsPerSecond)
	for {
		timeOutput := time.Now().Format("Jan 02 15:04:05")
		host := utils.RandomHost(cfg.Host, cfg.DomainName)
		tag := utils.RandomElement(cfg.Tags)
		pid := utils.RandomPid()
		message := utils.RandomElement(cfg.SampleLogs)
		level := utils.RandomElement(cfg.ParsedSyslogLevels)

		formattedMessage := fmt.Sprintf("%s %s %s[%d]: %s", timeOutput, host, tag, pid, message)
		fmt.Printf("[+] Sent: %s\n", formattedMessage)
		logger.Send(formattedMessage, level)

		time.Sleep(interval)
	}
}

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("[+] Sending messages to %s on port %d. Press Ctrl+C to stop.\n", cfg.Host, cfg.Port)

	logger, err := syslog.NewLogger(cfg.Host, cfg.Port)
	if err != nil {
		log.Fatalf("Failed to connect to syslog server: %v", err)
	}
	defer logger.Close()

	sendSyslogs(logger, cfg)
}

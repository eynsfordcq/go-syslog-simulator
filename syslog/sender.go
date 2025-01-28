package syslog

import (
	"fmt"
	"time"

	"github.com/eynsfordcq/go-syslog-simulator/utils"
)

func (l *Logger) SendSyslogs() {
	interval := time.Second / time.Duration(l.config.RequestsPerSecond)
	for {
		timeOutput := time.Now().Format("Jan 02 15:04:05")
		host := utils.RandomHost(l.config.Host, l.config.DomainName)
		tag := utils.RandomElement(l.config.Tags)
		pid := utils.RandomPid()
		message := utils.RandomElement(l.config.SampleLogs)
		level := utils.RandomElement(l.config.ParsedSyslogLevels)

		formattedMessage := fmt.Sprintf("%s %s %s[%d]: %s", timeOutput, host, tag, pid, message)
		fmt.Printf("[+] Sent: %s\n", formattedMessage)
		l.Send(formattedMessage, level)

		time.Sleep(interval)
	}
}

package syslog

import (
	"fmt"
	"time"

	"github.com/eynsfordcq/go-syslog-simulator/utils"
)

func (l *Logger) SendSyslogs() {
	interval := time.Second / time.Duration(l.config.RequestsPerSecond)
	for {
		level := utils.RandomElement(l.config.ParsedSyslogLevels)
		pri := int(l.config.ParsedFacility) + int(level)

		timeOutput := time.Now().Format("Jan 02 15:04:05")
		host := utils.RandomHost(l.config.Hostname, l.config.DomainName)
		tag := utils.RandomElement(l.config.Tags)
		pid := utils.RandomPid()
		message := utils.RandomElement(l.config.SampleLogs)

		formattedMessage := fmt.Sprintf(
			"<%d>%s %s %s[%d]: %s\n",
			pri,
			timeOutput,
			host,
			tag,
			pid,
			message,
		)

		fmt.Printf("[+] Sent: %s", formattedMessage)
		l.Send(formattedMessage)

		time.Sleep(interval)
	}
}

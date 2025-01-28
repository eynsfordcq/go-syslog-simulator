package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/eynsfordcq/go-syslog-simulator/config"
	"github.com/eynsfordcq/go-syslog-simulator/syslog"
	"github.com/spf13/cobra"
)

var (
	host              string
	requestsPerSecond int
	port              int
	configFile        string
)

var rootCmd = &cobra.Command{
	Use:   "go-syslog-simulator",
	Short: "Simulate syslog messages",
	Long:  `A tool to simulate syslog messages to a specified syslog server.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(configFile)
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		cfg.Host = host
		cfg.Port = port
		cfg.RequestsPerSecond = requestsPerSecond

		fmt.Printf("[+] Sending messages to %s on port %d. Press Ctrl+C to stop.\n", cfg.Host, cfg.Port)

		logger, err := syslog.NewLogger(cfg)
		if err != nil {
			log.Fatalf("Failed to connect to syslog server: %v", err)
		}
		defer logger.Close()

		logger.SendSyslogs()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&host, "host", "H", "127.0.0.1", "Syslog server host")
	rootCmd.Flags().IntVarP(&requestsPerSecond, "requests", "r", 1, "Requests per second")
	rootCmd.Flags().IntVarP(&port, "port", "p", 514, "Syslog server port")
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "Path to config file")
}

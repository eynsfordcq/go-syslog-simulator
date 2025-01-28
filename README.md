# Go Syslog Simulator

A tool to simulate syslog messages. This program can be used for testing and validation of syslog servers by generating a configurable number of syslog messages.


## Usage

You can run the program via the command line:

```bash
A tool to simulate syslog messages to a specified syslog server.

Usage:
  go-syslog-simulator [flags]

Flags:
  -c, --config string   Path to config file (default "config.json")
  -h, --help            help for go-syslog-simulator
  -H, --host string     Syslog server host (default "127.0.0.1")
  -p, --port int        Syslog server port (default 514)
  -r, --requests int    Requests per second (default 1)
```

## Configuration

Make a copy of `config.json.example` and make necessary adjustments. 

```bash
cp config.json.example config.json
```
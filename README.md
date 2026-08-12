# LogSage

LogSage is a Go-based server log analyser that reads log files, parses individual log entries, categorises log levels, and generates a summary report.

## Features

- Reads server logs from a file
- Parses raw log lines into structured `LogEntry` objects
- Supports `ERROR`, `INFO`, and `WARNING` log levels
- Detects unknown log levels
- Handles malformed log lines without stopping the entire analysis
- Generates a clean summary report
- Uses a modular package structure

## Project Structure

```text
LogSage/
│
├── analyser/
│   └── analyser.go
│
├── reader/
│   └── reader.go
│
├── report/
│   └── report.go
│
├── models/
│
├── utils/
│
├── main.go
├── server.log
├── go.mod
├── .gitignore
└── README.md
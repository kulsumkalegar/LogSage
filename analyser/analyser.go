package analyser

import (
	"errors"
	"strings"
)

type LogEntry struct {
	TimeStamp string
	Level     string
	Message   string
}

func parseLine(line string) (LogEntry, error) {
	word := strings.Fields(line)
	if len(word) < 3 {
		return LogEntry{}, errors.New("invalid log line")
	} else {
		logparser := LogEntry{
			TimeStamp: strings.Join(word[0:2], " "),
			Level:     word[2],
			Message:   strings.Join(word[3:], " "),
		}
		return logparser, nil
	}

}

func Analyse(lines []string) ([]LogEntry, int, int, int, int, error) {
	errorcount := 0
	infocount := 0
	warningcount := 0
	unknowncount := 0
	var entries []LogEntry
	if len(lines) == 0 {
		return entries,
			errorcount, infocount, warningcount, unknowncount, errors.New("Empty log file")
	} else {
		for _, line := range lines {
			entry, err := parseLine(line)
			if err != nil {
				continue
			}
			if entry.Level == "ERROR" {
				errorcount += 1

			} else if entry.Level == "INFO" {
				infocount += 1
			} else if entry.Level == "WARNING" {
				warningcount += 1
			} else {
				unknowncount++
			}
			entries = append(entries, entry)
		}

	}

	return entries, errorcount, infocount, warningcount, unknowncount, nil

}

package cloudFunction

import (
	"encoding/json"
	"fmt"
	"log"
)

	type logEntry struct {

		Message  string `json:"message"`
		Severity string `json:"severity,omitempty"`
		Trace    string `json:"logging.googleapis.com/trace,omitempty"`

		// Stackdriver Log Viewer allows filtering and display of this as `jsonPayload.component`.
		Component string `json:"component,omitempty"`
	}

	type Log struct {

		entry logEntry
	}

	func (l *Log) Debug(msg... interface{}) {

		l.doLog("DEBUG", fmt.Sprintln(msg...))
		fmt.Println(l.string())
	}

	func (l *Log) Info(msg... interface{}) {

		l.doLog("INFO", fmt.Sprintln(msg...))
		fmt.Println(l.string())
	}

	func (l *Log) Notice(msg... interface{}) {

		l.doLog("NOTICE", fmt.Sprintln(msg...))
		fmt.Println(l.string())
	}

	func (l *Log) Warning(msg... interface{}) {

		l.doLog("WARNING", fmt.Sprintln(msg...))
		fmt.Println(l.string())
	}

	func (l *Log) Error(msg... interface{}) {

		l.doLog("ERROR", fmt.Sprintln(msg...))
		fmt.Println(l.string())
	}

	func (l *Log) Critical(msg... interface{}) {

		l.doLog("CRITICAL", fmt.Sprintln(msg...))
		fmt.Println(l.string())
	}

	func (l *Log) Alert(msg... interface{}) {

		l.doLog("ALERT", fmt.Sprintln(msg...))
		fmt.Println(l.string())
	}

	func (l *Log) Emergency(msg... interface{}) {

		l.doLog("EMERGENCY", fmt.Sprintln(msg...))
		fmt.Println(l.string())
	}

	func (l *Log) doLog(level, msg string) {

		l.entry.Message = msg
		l.entry.Severity = level
	}

	func (l *Log) string() string {

		if l.entry.Severity == "" {
			l.entry.Severity = "INFO"
		}
		out, err := json.Marshal(l.entry)
		if err != nil {
			log.Printf("json.Marshal: %v", err)
		}
		return string(out)
	}

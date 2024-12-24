package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nadirakdag/finance-tracker/internal/config"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct {
	level     LogLevel
	isJSON    bool
	stdLogger *log.Logger
}

type LogEntry struct {
	Timestamp string      `json:"timestamp"`
	Level     string      `json:"level"`
	Message   string      `json:"message"`
	Fields    interface{} `json:"fields,omitempty"`
}

func NewLogger(cfg *config.LoggingConfig) *Logger {
	level := parseLevel(cfg.Level)
	isJSON := cfg.Format == "json"

	return &Logger{
		level:     level,
		isJSON:    isJSON,
		stdLogger: log.New(os.Stdout, "", 0), // No prefix or flags as we'll handle them
	}
}

func parseLevel(level string) LogLevel {
	switch level {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	default:
		return INFO
	}
}

func (l *Logger) log(level LogLevel, msg string, fields ...interface{}) {
	if level < l.level {
		return
	}

	if l.isJSON {
		l.logJSON(level, msg, fields...)
	} else {
		l.logText(level, msg, fields...)
	}
}

func (l *Logger) logJSON(level LogLevel, msg string, fields ...interface{}) {
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level.String(),
		Message:   msg,
	}

	if len(fields) > 0 {
		fieldMap := make(map[string]interface{})
		for i := 0; i < len(fields); i += 2 {
			if i+1 < len(fields) {
				fieldMap[fmt.Sprint(fields[i])] = fields[i+1]
			}
		}
		entry.Fields = fieldMap
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		l.stdLogger.Printf("Error marshaling log entry: %v", err)
		return
	}

	l.stdLogger.Println(string(jsonData))
}

func (l *Logger) logText(level LogLevel, msg string, fields ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelStr := level.String()

	// Format fields as key=value pairs
	var fieldStr string
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			fieldStr += fmt.Sprintf(" %v=%v", fields[i], fields[i+1])
		}
	}

	l.stdLogger.Printf("%s [%s] %s%s", timestamp, levelStr, msg, fieldStr)
}

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func (l *Logger) Debug(msg string, fields ...interface{}) {
	l.log(DEBUG, msg, fields...)
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	l.log(INFO, msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.log(WARN, msg, fields...)
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	l.log(ERROR, msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...interface{}) {
	l.log(FATAL, msg, fields...)
	os.Exit(1)
}

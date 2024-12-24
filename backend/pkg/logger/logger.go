package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	warnLogger  *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLogger:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string, keyvals ...interface{}) {
	l.log(l.infoLogger, msg, keyvals...)
}

func (l *Logger) Error(msg string, keyvals ...interface{}) {
	l.log(l.errorLogger, msg, keyvals...)
}

func (l *Logger) Warn(msg string, keyvals ...interface{}) {
	l.log(l.warnLogger, msg, keyvals...)
}

func (l *Logger) Fatal(msg string, keyvals ...interface{}) {
	l.log(l.errorLogger, msg, keyvals...)
	os.Exit(1)
}

func (l *Logger) log(logger *log.Logger, msg string, keyvals ...interface{}) {
	args := make([]interface{}, 0, len(keyvals)+1)
	args = append(args, msg)
	args = append(args, keyvals...)
	logger.Println(args...)
}

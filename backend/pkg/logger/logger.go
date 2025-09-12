package logger

import (
	"log"
	"os"
)

// Logger interface for structured logging
type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Debug(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
}

// StandardLogger implements Logger interface using standard log package
type StandardLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
	warnLogger  *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() Logger {
	return &StandardLogger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Lshortfile),
		warnLogger:  log.New(os.Stdout, "WARN: ", log.LstdFlags|log.Lshortfile),
	}
}

// Info logs info level messages
func (l *StandardLogger) Info(msg string, fields ...interface{}) {
	l.infoLogger.Printf(msg, fields...)
}

// Error logs error level messages
func (l *StandardLogger) Error(msg string, fields ...interface{}) {
	l.errorLogger.Printf(msg, fields...)
}

// Debug logs debug level messages
func (l *StandardLogger) Debug(msg string, fields ...interface{}) {
	l.debugLogger.Printf(msg, fields...)
}

// Warn logs warning level messages
func (l *StandardLogger) Warn(msg string, fields ...interface{}) {
	l.warnLogger.Printf(msg, fields...)
}

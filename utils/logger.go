package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevel defines the types of log levels
const (
	INFO    = "INFO"
	SUCCESS = "SUCCESS"
	ERROR   = "ERROR"
	FATAL   = "FATAL"
)

// ANSI color codes for log levels
const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[34m"
	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
	colorPurple = "\033[35m"
)

// Logger is a custom logger with color support
type Logger struct {
	logger *log.Logger
}

// NewLogger creates a new instance of Logger
func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", 0),
	}
}

// logMessage prints a log message with the given level and color
func (l *Logger) logMessage(level string, color string, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formattedMessage := fmt.Sprintf("%s[%s] [%s] %s%s", color, timestamp, level, message, colorReset)
	l.logger.Println(formattedMessage)
}

// Info logs an informational message
func (l *Logger) Info(message string) {
	l.logMessage(INFO, colorBlue, message)
}

// Success logs a success message
func (l *Logger) Success(message string) {
	l.logMessage(SUCCESS, colorGreen, message)
}

// Error logs an error message
func (l *Logger) Error(message string) {
	l.logMessage(ERROR, colorRed, message)
}

// Fatal logs a fatal message and exits the program
func (l *Logger) Fatal(message string) {
	l.logMessage(FATAL, colorPurple, message)
	os.Exit(1)
}

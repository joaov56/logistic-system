package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger represents a simple logger
type Logger struct {
	*log.Logger
}

// New creates a new logger
func New() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", 0),
	}
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	l.log("INFO", format, v...)
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	l.log("ERROR", format, v...)
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	l.log("DEBUG", format, v...)
}

// log logs a message with the given level
func (l *Logger) log(level, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.Printf("[%s] %s %s", level, time.Now().Format("2006-01-02 15:04:05"), msg)
} 
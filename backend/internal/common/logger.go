package utils

import (
	"log"
)

type Level string

const (
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO"
)

type Logger struct {
	LogLevel Level
}

func NewLogger() *Logger {
	return &Logger{
		LogLevel: LevelInfo,
	}
}

func (l *Logger) SetLogLevel(level Level) {
	if level != LevelDebug && level != LevelInfo {
		l.LogLevel = LevelInfo
		log.Printf("invalid log level: %s. defaulting to INFO.\n", level)
	}
	l.LogLevel = level
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.LogLevel == LevelDebug {
		log.Printf("🚧 DEBUG: "+msg, args...)
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	log.Printf("  INFO: "+msg, args...)
}
func (l *Logger) Error(msg string, args ...interface{}) {
	log.Printf("❌ ERROR: "+msg, args...)
}
func (l *Logger) Fatal(msg string, args ...interface{}) {
	log.Fatalf("❌ FATAL: "+msg, args...)
}

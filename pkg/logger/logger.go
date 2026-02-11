package logger

import "go.uber.org/zap"

// Global logger instance
var globalLogger *zap.Logger

// NewLogger creates and returns a new structured logger instance
func NewLogger() (*zap.Logger, error) {
	var err error
	globalLogger, err = zap.NewProduction()
	return globalLogger, err
}

// Console logs a message to console with optional fields
// Usage: Console("user created", zap.String("userID", "123"), zap.Int("status", 201))
func Console(msg string, fields ...zap.Field) {
	if globalLogger != nil {
		globalLogger.Info(msg, fields...)
	}
}

// ConsoleError logs an error message to console with optional fields
func ConsoleError(msg string, fields ...zap.Field) {
	if globalLogger != nil {
		globalLogger.Error(msg, fields...)
	}
}

// ConsoleWarn logs a warning message to console with optional fields
func ConsoleWarn(msg string, fields ...zap.Field) {
	if globalLogger != nil {
		globalLogger.Warn(msg, fields...)
	}
}

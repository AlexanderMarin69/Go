package logger

import "go.uber.org/zap"

// NewLogger creates and returns a new structured logger instance
func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

/*
Package Name: lib
File Name: logger.go
Abstract: The logger used for logging data.
*/
package lib

import (
	"os"

	"github.com/withmandala/go-log"
)

// ======== TYPES ========

// loggerInterface represents the logging functionality used in the code.
// This is done to enable mocking the logger in tests.
type Logger interface {
	Info(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
}

// NewHandler returns a new gin router
func GetLogger() Logger {
	logger := log.New(os.Stderr)
	return logger
}
